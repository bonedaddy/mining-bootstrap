package ethermine_test

import (
	"errors"
	"testing"

	"github.com/RTradeLtd/mining-bootstrap/earnings-report/reports/ethermine"
)

const (
	testMiner = "0x7509b8b9c013AA1280cC8Bf5Ec5908856faCc5a1"
)

func TestReport(t *testing.T) {
	manager, err := ethermine.GenerateReportManagerFromFile("", false)
	if err != nil {
		t.Fatal(err)
	}
	payouts, err := manager.GetPayouts(testMiner)
	if err != nil {
		t.Fatal(err)
	}
	if len(*payouts) == 0 {
		err := errors.New("no payouts retrieved")
		t.Fatal(err)
	}
}
