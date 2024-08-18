package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/xuri/excelize/v2"
)

const (
	basedir = "."
)

type Response struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hourly    Hourly  `json:"hourly"`
}

type Hourly struct {
	Temperature2m []float64 `json:"temperature_2m"`
	Time          []string  `json:"time"`
}

func main() {
	// Define the API URL
	url := "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&hourly=temperature_2m"

	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return
	}

	// Parse the JSON response
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing the JSON response:", err)
		return
	}

	converToExcel(response)
	convertToPdf(response)
}

func converToExcel(response Response) {
	// Create a new Excel file
	namesFile := excelize.NewFile()
	namesFileSheet1Index, _ := namesFile.NewSheet("Sheet1")
	namesFile.SetActiveSheet(namesFileSheet1Index)

	// Write latitude and longitude
	_ = namesFile.SetCellValue("Sheet1", "A1", "Latitude")
	_ = namesFile.SetCellValue("Sheet1", "A2", "Longitude")
	_ = namesFile.SetCellValue("Sheet1", "B1", response.Latitude)
	_ = namesFile.SetCellValue("Sheet1", "B2", response.Longitude)

	// title column
	_ = namesFile.SetCellValue("Sheet1", "A4", "Tanggal & Jam")
	_ = namesFile.SetCellValue("Sheet1", "B4", "Suhu")

	// Write temperature values
	for i, temp := range response.Hourly.Temperature2m {
		cell := fmt.Sprintf("B%d", i+5)
		_ = namesFile.SetCellValue("Sheet1", cell, temp)
	}

	// Write time values
	for i, temp := range response.Hourly.Time {
		cell := fmt.Sprintf("A%d", i+5)
		_ = namesFile.SetCellValue("Sheet1", cell, temp)
	}

	// Save the Excel file
	if err := namesFile.SaveAs("weather_data.xlsx"); err != nil {
		fmt.Println("Error saving the Excel file:", err)
	}
}

func convertToPdf(response Response) {
	template, err := template.ParseFiles(basedir + "/static/template.html")
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	data := Response{
		Latitude:  response.Latitude,
		Longitude: response.Longitude,
	}

	// Write temperature values
	for _, temp := range response.Hourly.Temperature2m {
		data.Hourly.Temperature2m = append(data.Hourly.Temperature2m, temp)
	}

	// // Write time values
	for _, temp := range response.Hourly.Time {
		data.Hourly.Time = append(data.Hourly.Time, temp)
	}

	// Save import data to template
	if err := template.Execute(buf, data); err != nil {
		log.Println(err)
	}

	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}
	// Create a new input page from an URL
	page := wkhtmltopdf.NewPageReader(strings.NewReader(buf.String()))

	// Add to document
	pdfg.AddPage(page)

	// Create PDF document in internal buffer
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile("weather_data.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
}
