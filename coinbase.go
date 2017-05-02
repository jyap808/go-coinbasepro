// Package Coinbase is an implementation of the Coinbase API in Golang.
package coinbase

import (
	"encoding/json"
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

// coinbase represent a coinbase client
type Coinbase struct {
	client *client
}

// https://api.coinbase.com/v2/prices/BTC-USD/spot

// GetTicker is used to get the current ticker values for a market.
func (b *Coinbase) GetTicker(market string) (ticker Ticker, err error) {
	r, err := b.client.do("GET", "pubticker/"+strings.ToLower(market), "", false)
	if err != nil {
		return
	}
	if err = json.Unmarshal(r, &ticker); err != nil {
		return
	}
	return
}
