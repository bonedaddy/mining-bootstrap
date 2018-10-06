package utils

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"time"

	"github.com/RTradeLtd/mining-bootstrap/earnings-report/reports/types"
)

const (
	// USDAPI is the URL We use to query for USD->CAD conversion
	USDAPI = "https://free.currencyconverterapi.com/api/v5/convert?q=USD_CAD&compact=y"
)

// ParseETHUSD is used to retrieve the latest ETH->USD price
func ParseETHUSD() (float64, error) {
	ethUSD, err := types.RetrieveEthUsdPrice()
	if err != nil {
		return 0, err
	}
	return ethUSD, nil
}

// ParseUSDCAD is used to retrieve the latest USD -> CAD conversion ratio
func ParseUSDCAD() (float64, error) {
	client := http.DefaultClient
	client.Timeout = time.Minute
	resp, err := client.Get(USDAPI)
	if err != nil {
		return 0, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var intf map[string]interface{}
	err = json.Unmarshal(respBytes, &intf)
	if err != nil {
		return 0, err
	}
	marshaled, err := json.Marshal(intf["USD_CAD"])
	if err != nil {
		return 0, err
	}
	var val types.USDResponse
	err = json.Unmarshal(marshaled, &val)
	if err != nil {
		return 0, err
	}
	return val.ExchangeRate, nil
}

// BaseWeiToBaseEth is used to convert a number from it's wei representation to it's eth representation
func BaseWeiToBaseEth(x float64) float64 {
	exp := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	xBig := big.NewFloat(x)
	floatExp := float64(exp.Int64())
	div := new(big.Float).Quo(xBig, big.NewFloat(floatExp))
	f, _ := div.Float64()
	return f
}
