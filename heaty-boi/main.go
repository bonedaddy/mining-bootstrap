package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

/*
rtrade@condo2-floor2-rig4:~$ nvidia-smi --query-gpu=temperature.gpu --format=csv,noheader,nounits -q -i 1
43
*/

const (
	argCount = 2
)

func main() {
	if len(os.Args) > argCount || len(os.Args) < argCount {
		log.Fatal("bad arg count")
	}
	number := os.Args[1]
	for {
		records, err := getGPUTemp(number)
		if err != nil {
			log.Fatal(err)
		}
		for _, row := range records {
			index := row[0]
			temp := row[1]
			fmt.Printf("GPU %s TEMP %sC\n", index, temp)
			tempInt, err := strconv.ParseInt(temp, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			if tempInt >= 80 {
				fmt.Println("gpu temp above 80, stopping miner")
				_, err := exec.Command(
					"systemctl",
					"stop",
					"miner",
				).Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("miner stopped, waiting for chips to cool")
				fmt.Println("sleeping for 2 minutes")
				time.Sleep(time.Minute * 2)
				fmt.Println("sleep over, resuming miner")
				_, err = exec.Command(
					"systemctl",
					"start",
					"miner",
				).Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("miner resumed")
				break
			}
		}
	}
}

func getGPUTemp(gpuNumber string) ([][]string, error) {
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
