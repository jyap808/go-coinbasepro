package coinbasepro

import (
	"encoding/json"
)

type Stats24Hour struct {
	Volume json.Number `json:"volume"`
	Last   json.Number `json:"last"`
}

type Stat struct {
	Stats24Hour Stats24Hour `json:"stats_24hour"`
}
