package types

// USDResponse is used to format the resposne from our USD->CAD api
type USDResponse struct {
	ExchangeRate float64 `json:"val"`
}
