package coinbasepro

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	apiURL            = "https://api.exchange.coinbase.com/" // Coinbase Pro API endpoint
	httpClientTimeout = 30                                   // HTTP client timeout
)

// Client holds information necessary to make a request to your API
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// Option is a functional option for configuring the API client
type Option func(*Client) error

// BaseURL allows overriding of API client baseURL for testing
func BaseURL(baseURL string) Option {
	return func(c *Client) error {
		c.baseURL = baseURL
		return nil
	}
}

// New creates a new API client
func New(opts ...Option) *Client {
	client := &Client{
		baseURL: apiURL,
		httpClient: &http.Client{
			Timeout: time.Second * httpClientTimeout,
		},
	}

	if err := client.parseOptions(opts...); err != nil {
		return nil
	}

	return client
}

// parseOptions parses the supplied options functions and returns a configured
// *Client instance
func (c *Client) parseOptions(opts ...Option) error {
	// Range over each options function and apply it to our API type to
	// configure it. Options functions are applied in order, with any
	// conflicting options overriding earlier calls.
	for _, option := range opts {
		err := option(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// doTimeoutRequest do a HTTP request with timeout
func (c *Client) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	// Do the request in the background so we can check the timeout
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := c.httpClient.Do(req)
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("timeout on reading data from Coinbase API")
	}
}

// do prepare and process HTTP request to Coinbase API
func (c *Client) do(method string, resource string, payload string, authNeeded bool) (response []byte, err error) {
	connectTimer := time.NewTimer(httpClientTimeout * time.Second)

	var rawurl string
	if strings.HasPrefix(resource, "http") {
		rawurl = resource
	} else {
		rawurl = fmt.Sprintf("%s%s", apiURL, resource)
	}

	req, err := http.NewRequest(method, rawurl, strings.NewReader(payload))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")

	resp, err := c.doTimeoutRequest(connectTimer, req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return response, err
}
