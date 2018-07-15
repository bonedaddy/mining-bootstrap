package reports_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/RTradeLtd/mining-bootstrap/src/reports"
)

func TestManager(t *testing.T) {
	manager, err := reports.GenerateReportManagerFromFile()
	if err != nil {
		t.Fatal(err)
	}
	// format the URL
	manager.FormatURL("getdashboarddata")
	// make the request
	resp, err := http.Get(manager.Config.URL)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
