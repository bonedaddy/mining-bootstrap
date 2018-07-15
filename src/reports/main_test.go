package reports_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/RTradeLtd/mining-bootstrap/src/reports"
	"github.com/RTradeLtd/mining-bootstrap/src/types"
)

func TestManager(t *testing.T) {
	manager, err := reports.GenerateReportManagerFromFile()
	if err != nil {
		t.Fatal(err)
	}
	// format the URL
	manager.FormatURL("getdashboarddata")
	fmt.Println(manager.Config.URL)
	// make the request
	resp, err := http.Get(manager.Config.URL)
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var intf map[string]interface{}
	var data types.GetDashboardData
	err = json.Unmarshal(body, &intf)
	if err != nil {
		t.Fatal(err)
	}
	marshaled, err := json.Marshal(intf["getdashboarddata"])
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(marshaled, &data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v\n", data.Data["recent_credits"])
}
