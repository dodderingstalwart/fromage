package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func main() {
	fmt.Println("Fromage")

	ctx := context.Background()

        // Check for json credentials file 
	cred, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read credentials file: %v", err)
	}

        // Check if credentials file is correct
	config, err := google.JWTConfigFromJSON(cred, sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// The ID of the spreadsheet to retrieve data from.
	spreadsheetId := os.Getenv("GSHEETID")
	// The range of cells to read from the spreadsheet.
	readRange := "Sheet1!A1:K22"

	// Retrieve data from the specified range
	retrieve, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(retrieve.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Data from sheet:")
		for _, row := range retrieve.Values {
			fmt.Println(row)
		}
	}
}
