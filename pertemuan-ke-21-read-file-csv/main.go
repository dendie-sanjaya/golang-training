package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

const (
	basedir  = "."
	fileName = "data.csv"
)

func main() {
	csvFile, err := os.Open(basedir + "/" + fileName)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	reportFile, _ := os.Create(basedir + "/report.log")
	reportFileWriter := bufio.NewWriter(reportFile)
	_ = reportFileWriter.Flush()

	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatalf("failed reading csv file: %s", err)
	}
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		name, address := line[0], line[1]
		report := fmt.Sprintf("Konten baris ke %d: name: %s address: %s \n", i, name, address)
		_, _ = fmt.Fprintf(reportFileWriter, report)
		_ = reportFileWriter.Flush()
	}
}
