package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	fmt.Println("Fromage")
	ctx := context.Background()

	// Initialize the Sheets API client
	srv, err := sheets.NewService(ctx, option.WithCredentialsFile("credentials.json"))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// The ID of the spreadsheet to retrieve data from.
	spreadsheetId := os.Getenv("GSHEETID")

	// The range of cells to read from the spreadsheet.
	readRange := "Sheet1!A1:D10"

	// Retrieve data from the specified range
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data from sheet:")
		for _, row := range resp.Values {
			fmt.Println(row)
		}
	}
}
