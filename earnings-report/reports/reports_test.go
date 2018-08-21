package reports_test

import (
	"fmt"
	"testing"

	"github.com/RTradeLtd/mining-bootstrap/src/reports"
)

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
	t.Skip()
	manager, err := reports.GenerateReportManagerFromFile("")
	if err != nil {
		t.Fatal(err)
	}

	err = manager.CreateReportAndSend("24hour_credit")
	if err != nil {
		t.Fatal(err)
	}
}
