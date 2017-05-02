package coinbase

type Price struct {
	BTC      AmountCurrency `json:"btc"`
	Subtotal AmountCurrency `json:"subtotal"`
	Fees     interface{}    `json:"fees"`
	Total    AmountCurrency `json:"total"`
	Amount   float64        `json:"amount,string"`
	Currency string         `json:"currency"`
	Error    string         `json:"error"`
}

type AmountCurrency struct {
	Amount   float64 `json:"amount,string"`
	Currency string  `json:"currency"`
}
