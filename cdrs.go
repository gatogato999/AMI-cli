package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func printCDR() {
	file, err := os.Open("/var/log/asterisk/cdr-custom/Simple.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}

		fmt.Printf("source: %s, dest : %s, channel : %s, duration: %s, billsec: %s\n",
			record[0], record[1], record[2], record[3], record[4])
	}
}
