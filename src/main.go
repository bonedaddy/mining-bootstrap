package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RTradeLtd/mining-bootstrap/src/reports"
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
	fmt.Println("generating 24 hour earnings report")
	err = m.CreateReportAndSend("24hour_credit")
	if err != nil {
		log.Fatal(err)
	}
}
