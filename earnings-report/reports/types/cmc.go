package types

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Response used to hold response data from cmc
type Response struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Symbol             string `json:"symbol"`
	Rank               string `json:"rank"`
	PriceUsd           string `json:"price_usd"`
	PriceBtc           string `json:"price_btc"`
	TwentyFourHrVolume string `json:"24h_volume_usd"`
	MarketCapUsd       string `json:"market_cap_usd"`
	AvailableSupply    string `json:"available_supply"`
	TotalSupply        string `json:"total_supply"`
	MaxSupply          string `json:"null"`
	PercentChange1h    string `json:"percent_change_1h"`
	PercentChange24h   string `json:"percent_change_24h"`
	PercentChange7d    string `json:"percent_change_7d"`
	LastUpdate         string `json:"last_updated"`
}

// RetrieveEthUsdPrice is used to retrieve the ETH->USD price
func RetrieveEthUsdPrice() (float64, error) {
	client := http.DefaultClient
	client.Timeout = time.Minute
	response, err := client.Get("https://api.coinmarketcap.com/v1/ticker/ethereum/")
	if err != nil {
		return float64(0), err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return float64(0), err
	}
	var decode []Response
	err = json.Unmarshal(body, &decode)
	if err != nil {
		return float64(0), err
	}

	// TODO: add error handling
	f, err := strconv.ParseFloat(decode[0].PriceUsd, 64)
	if err != nil {
		return float64(0), err
	}

	return f, nil
}
