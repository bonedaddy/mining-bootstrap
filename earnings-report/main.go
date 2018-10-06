package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/RTradeLtd/mining-bootstrap/earnings-report/reports/mph"
)

func main() {

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		err := errors.New("CONFIG_PATH env variable not set, using default config path")
		log.Fatal(err)
	}
	runMode := os.Getenv("RUN_MODE")
	if runMode == "" {
		log.Fatal("RUN_MODE is not set")
	}
	switch runMode {
	case "report", "report-save":
		break
	default:
		err := errors.New("RUN_MODE must be 'report' or 'report-save'")
		log.Fatal(err)
	}
	if len(os.Args) > 2 || len(os.Args) < 2 {
		err := errors.New("invalid argument provided must be: 'mph'")
		log.Fatal(err)
	}
	pool := os.Args[1]
	switch pool {
	case "mph":
		if runMode == "report-save" {
			m, err := mph.GenerateReportManagerFromFile(configPath, true)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("generating 24 hour earnings report and saving to DB")
			err = m.GetRecentCredits24HoursAndSave()
			if err != nil {
				log.Fatal(err)
			}
		} else if runMode == "report" {
			m, err := mph.GenerateReportManagerFromFile(configPath, false)
			if err != nil {
				log.Fatal(err)
			}
			credits, err := m.GetRecentCredits24Hours()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Credits received:\n%+v\n", credits)
		}

	}

}
