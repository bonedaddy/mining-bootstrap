package types

// MiningPoolHubAPIResponse is used to format the response from mining pool hub
type MiningPoolHubAPIResponse struct {
	Version string                 `json:"version"`
	Runtime float64                `json:"runtime"`
	Data    map[string]interface{} `json:"data"`
}

// RecentCredits is used to format the response from mining pool hub's recent credits call
type RecentCredits struct {
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
}

// USDResponse is used to format the resposne from our USD->CAD api
type USDResponse struct {
	ExchangeRate float64 `json:"val"`
}
