package reports

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/RTradeLtd/mining-bootstrap/src/reports/types"
)

func ParseETHUSD() (float64, error) {
	ethUSD, err := types.RetrieveEthUsdPrice()
	if err != nil {
		return 0, err
	}
	return ethUSD, nil
}

func ParseUSDCAD() (float64, error) {
	resp, err := http.Get(USDAPI)
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
