package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need just one filename!")
		os.Exit(-1)
	}

	filename := os.Args[1]
	_, err := os.Stat(filename)
	if err == nil {
		fmt.Printf("File %s already exists.\n", filename)
		os.Exit(1)
	}

	output, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer output.Close()

	inputData := [][]string{{"M", "T", "I."}, {"D", "T", "I."}, {"M", "T", "D."}, {"V", "T", "D."}, {"A", "T", "D."}}
	writer := csv.NewWriter(output)
	for _, record := range inputData {
		err := writer.Write(record)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}
	writer.Flush()

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	allRecords, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, rec := range allRecords {
		fmt.Printf("%s:%s:%s\n", rec[0], rec[1], rec[2])
	}
}
