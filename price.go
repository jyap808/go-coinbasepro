package coinbase

type AmountCurrency struct {
	Amount   float64 `json:"amount,string"`
	Currency string  `json:"currency"`
}
