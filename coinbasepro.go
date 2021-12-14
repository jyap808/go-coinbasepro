// Package Coinbase Pro is an implementation of the Coinbase Pro API in Golang.
package coinbasepro

import (
	"encoding/json"
)

// https://api.exchange.coinbase.com/currencies
// Get all known currencies
func (b *Client) GetCurrencies() (currencies []Currency, err error) {
	r, err := b.do("GET", "currencies", "", false)
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
func (b *Client) GetStats() (stats map[string]Stat, err error) {
	r, err := b.do("GET", "products/stats", "", false)
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
