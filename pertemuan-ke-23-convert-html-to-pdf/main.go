package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

const (
	basedir      = "."
	htmlTemplate = "static/template.html"
	pdfOutput    = "static/output.pdf"
)

type Invoice struct {
	FullName      string
	Name          string
	Email         string
	Phone         string
	InvoiceNumber string
	SubTotal      string
}

func main() {
	template, err := template.ParseFiles(basedir + "/" + htmlTemplate)
	if err != nil {
		log.Fatal(err)
	}
	buf := new(bytes.Buffer)
	data := Invoice{
		FullName:      "Fahmi Hidayat",
		Name:          "Ibam",
		Email:         "ibrahimker@gmail.com",
		Phone:         "+1221412",
		InvoiceNumber: "INC1049",
		SubTotal:      "Rp.10.000.000,-",
	}
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
	err = pdfg.WriteFile(basedir + "/" + pdfOutput)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
	// Output: Done
}
