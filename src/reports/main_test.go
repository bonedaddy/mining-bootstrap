package reports_test

import (
	"fmt"
	"testing"

	"github.com/RTradeLtd/mining-bootstrap/src/reports"
)

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
