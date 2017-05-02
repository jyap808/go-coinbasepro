// Package Coinbase is an implementation of the Coinbase API in Golang.
package coinbase

import (
	"encoding/json"
	"errors"
	"strings"
)

const (
	API_BASE                   = "https://api.coinbase.com/" // Coinbase API endpoint
	API_VERSION                = "v2"                        // Coinbase API version
	DEFAULT_HTTPCLIENT_TIMEOUT = 30                          // HTTP client timeout
)

// New return a instanciate coinbase struct
func New(apiKey, apiSecret string) *Coinbase {
	client := NewClient(apiKey, apiSecret)
	return &Coinbase{client}
}

// handleErr gets JSON response from Coinbase API en deal with error
func handleErr(r jsonResponse) error {
	if r.Data == nil {
		return errors.New(r.Warnings)
	}
	return nil
}

// coinbase represent a coinbase client
type Coinbase struct {
	client *client
}

// https://api.coinbase.com/v2/prices/BTC-USD/spot

// GetTicker is used to get the current ticker values for a market.
func (b *Coinbase) GetPrices(market string) (amountcurrency AmountCurrency, err error) {
	r, err := b.client.do("GET", "prices/"+strings.ToLower(market)+"/spot", "", false)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Data, &amountcurrency)
	return

	//	if err != nil {
	//		return
	//	}
	//	if err = json.Unmarshal(r, &amountcurrency); err != nil {
	//		return
	//	}
	//	return
}
