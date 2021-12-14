// Package Coinbase Pro is an implementation of the Coinbase Pro API in Golang.
package coinbasepro

import (
	"encoding/json"
)

const (
	API_BASE                   = "https://api.exchange.coinbase.com/" // Coinbase Pro API endpoint
	DEFAULT_HTTPCLIENT_TIMEOUT = 30                                   // HTTP client timeout
)

// New return a instanciate coinbase struct
func New() *CoinbasePro {
	client := NewClient()
	return &CoinbasePro{client}
}

// coinbase represent a coinbase client
type CoinbasePro struct {
	client *client
}

// https://api.exchange.coinbase.com/currencies
// Get all known currencies
func (b *CoinbasePro) GetCurrencies() (currencies []Currency, err error) {
	r, err := b.client.do("GET", "currencies", "", false)
	if err != nil {
		return
	}

	if err = json.Unmarshal(r, &currencies); err != nil {
		return
	}
	if err != nil {
		return
	}
	return
}

// https://api.exchange.coinbase.com/products/stats
func (b *CoinbasePro) GetStats() (stats map[string]Stat, err error) {
	r, err := b.client.do("GET", "products/stats", "", false)
	if err != nil {
		return
	}

	if err = json.Unmarshal(r, &stats); err != nil {
		return
	}
	if err != nil {
		return
	}
	return
}
