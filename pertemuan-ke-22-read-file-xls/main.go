package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

const (
	basedir  = "."
	fileName = "data.xlsx"
)

func main() {
	excelFile, _ := excelize.OpenFile(basedir + "/" + fileName)
	rows, _ := excelFile.GetRows("Sheet1")
	var names []string
	for _, row := range rows {
		var name string
		for i, colCell := range row {
			fmt.Print(colCell + ",")
			if i == 1 {
				name += colCell
			}
			if i == 2 {
				name += colCell
			}
		}
		names = append(names, name)
		fmt.Println()
	}

	namesFile := excelize.NewFile()
	namesFileSheet1Index, _ := namesFile.NewSheet("Sheet1")
	namesFile.SetActiveSheet(namesFileSheet1Index)
	_ = namesFile.SetCellValue("Sheet1", "A1", "Name")
	for i, name := range names {
		_ = namesFile.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), name)
	}
	_ = namesFile.SaveAs(basedir + "/" + "names.xlsx")
}
