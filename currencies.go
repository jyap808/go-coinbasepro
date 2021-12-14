package coinbasepro

import "encoding/json"

type Currency struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Details Details `json:"details"`
}
type Details struct {
	NetworkConfirmations json.Number `json:"network_confirmations"`
}
