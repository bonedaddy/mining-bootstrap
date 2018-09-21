package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	for {
		records, err := getGPUTemp()
		if err != nil {
			log.Fatal(err)
		}
		if err := parseRecords(records); err != nil {
			log.Fatal(err)
		}
	}
}

func parseRecords(records [][]string) error {
	for _, row := range records {
		index := row[0]
		temp := row[1]
		fmt.Printf("GPU %s has a temp of %sC\n", index, temp)
		tempInt, err := strconv.ParseInt(temp, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		if tempInt >= 80 {
			fmt.Println("gpu temp above 80, stopping miner")
			if err := stopMiner(); err != nil {
				return err
			}
			fmt.Println("miner stopped, sleeping to cool chips")
			time.Sleep(time.Minute * 2)
			fmt.Println("sleep over, resumining miner")
			if err := startMiner(); err != nil {
				return err
			}
		}
	}
	return nil
}

func getGPUTemp() ([][]string, error) {
	out, err := exec.Command(
		"nvidia-smi",
		"--query-gpu=index,temperature.gpu",
		"--format=csv,noheader,nounits",
	).Output()
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(bytes.NewReader(out))
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()
	return records, nil
}

func stopMiner() error {
	_, err := exec.Command(
		"systemctl",
		"stop",
		"miner",
	).Output()
	return err
}

func startMiner() error {
	_, err := exec.Command(
		"systemctl",
		"start",
		"miner",
	).Output()
	return err
}
