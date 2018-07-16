package main

import (
	"log"

	"github.com/RTradeLtd/mining-bootstrap/src/reports"
)

func main() {
	m, err := reports.GenerateReportManagerFromFile("")
	if err != nil {
		log.Fatal(err)
	}
	err = m.CreateReportAndSend("24hour_credit")
	if err != nil {
		log.Fatal(err)
	}
}
