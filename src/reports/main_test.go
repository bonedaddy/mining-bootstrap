package reports_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/RTradeLtd/mining-bootstrap/src/reports"
	"github.com/RTradeLtd/mining-bootstrap/src/reports/types"
)

var usdAPI = "https://free.currencyconverterapi.com/api/v5/convert?q=USD_CAD&compact=y"

func TestCMC(t *testing.T) {
	ethUSD, err := types.RetrieveEthUsdPrice()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ethUSD)
}
func TestUSDAPI(t *testing.T) {
	resp, err := http.Get(usdAPI)
	if err != nil {
		t.Fatal(err)
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var intf map[string]interface{}
	err = json.Unmarshal(respBytes, &intf)
	if err != nil {
		t.Fatal(err)
	}
	marshaled, err := json.Marshal(intf["USD_CAD"])
	if err != nil {
		t.Fatal(err)
	}
	var val types.USDResponse
	err = json.Unmarshal(marshaled, &val)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", val)
}
func TestManager(t *testing.T) {
	manager, err := reports.GenerateReportManagerFromFile()
	if err != nil {
		t.Fatal(err)
	}
	credits, err := manager.GetRecentCredits()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", credits)
}
