package main

import (
	"encoding/csv"
	"log"
	"os"
)

func parseCsv(csvTitle string) [][]string {
	// open file
	f, err := os.Open(csvTitle)
	if err != nil {
		log.Fatal(err)
	}

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	w := csv.NewWriter(f)
	defer w.Flush()
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			print("csv reading error")
		}
	}(f)
	return data
}
