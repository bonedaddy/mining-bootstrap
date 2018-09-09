package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RTradeLtd/mining-bootstrap/earnings-report/reports"
)

func main() {

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		fmt.Println("CONFIG_PATH env variable not set, using default config path")
	}
	m, err := reports.GenerateReportManagerFromFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("generating 24 hour earnings report and saving to DB")
	err = m.GetRecentCredits24HoursAndSave()
	if err != nil {
		log.Fatal(err)
	}
}
