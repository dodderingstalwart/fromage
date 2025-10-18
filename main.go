package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	fmt.Println("Fromage")
	ctx := context.Background()

	b, err := google.CredentialsFromJSON(ctx, []byte(`{
		"type": "service_account",
	}`), sheets.SpreadsheetsReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse credentials from JSON: %v", err)
	}

	srv, err := sheets.NewService(ctx, option.WithCredentials(b))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := "your-spreadsheet-id"
	readRange := "Sheet1!A1:D10"
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
