package reports

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/RTradeLtd/mining-bootstrap/src/reports/types"
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
