package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/RTradeLtd/mining-bootstrap/earnings-report/reports/ethermine"

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
	case "report", "report-save", "payouts":
		break
	default:
		err := errors.New("RUN_MODE must be 'report' or 'report-save' or 'payouts'")
		log.Fatal(err)
	}
	if len(os.Args) > 3 || len(os.Args) < 2 {
		err := errors.New("invalid invocation. ./earnings-reports [mph|etheremine] <miner>")
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
			fmt.Println("generating and displaying 24 hour credit report")
			credits, err := m.GetRecentCredits24Hours()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Credits received:\n%+v\n", credits)
		}
	case "ethermine":
		if runMode == "payouts" {
			if len(os.Args) > 3 || len(os.Args) < 3 {
				err := errors.New("not enough arguments provided for ethermine payouts report")
				log.Fatal(err)
			}
			m, err := ethermine.GenerateReportManagerFromFile(configPath, false)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("generating and displaying payout data")
			miner := os.Args[2]
			payouts, err := m.GetPayouts(miner)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("displaying payouts")
			for _, v := range payouts {
				fmt.Println("-----------")
				m.PrettyPrintPayout(&v)
			}
		}
	}

}
