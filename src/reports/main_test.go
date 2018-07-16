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

func TestReport(t *testing.T) {
	manager, err := reports.GenerateReportManagerFromFile("")
	if err != nil {
		t.Fatal(err)
	}
	creds, err := manager.GetRecentCredits()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", creds)
}

func TestReport24Hour(t *testing.T) {
	manager, err := reports.GenerateReportManagerFromFile("")
	if err != nil {
		t.Fatal(err)
	}
	credit, err := manager.GetRecentCredits24Hours()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", credit)
}

func TestReportAndSend(t *testing.T) {
	manager, err := reports.GenerateReportManagerFromFile("")
	if err != nil {
		t.Fatal(err)
	}

	err = manager.CreateReportAndSend("24hour_credit")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCMC(t *testing.T) {
	t.Skip()
	ethUSD, err := types.RetrieveEthUsdPrice()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ethUSD)
}
func TestUSDAPI(t *testing.T) {
	t.Skip()
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
	t.Skip()
	manager, err := reports.GenerateReportManagerFromFile("")
	if err != nil {
		t.Fatal(err)
	}
	credits, err := manager.GetRecentCredits()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", credits)
}
