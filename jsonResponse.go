package coinbase

import (
	"encoding/json"
)

type jsonResponse struct {
	Data     json.RawMessage `json:"data"`
	Warnings string          `json:"warnings"`
}
