package main

import (
	"fmt"
	"os"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("recommendations.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	recommendationSvc := NewRecommendationService()

	rows, err := f.GetRows("Report Recommendations")
	for i, row := range rows {
		if i == 0 {
			continue
		}
		recommendationSvc.AddRecommendationFromSpreadsheetRow(row)
	}
	w := os.Stdout
	recommendationSvc.WriteJSON(w)
	fmt.Printf("Number of rows: %d\n", len(rows))
}
