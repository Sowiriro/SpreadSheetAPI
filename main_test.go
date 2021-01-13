package main

import (
	"fmt"
	"log"
	"testing"

	"google.golang.org/api/sheets/v4"
)

func TestReadSheet(t *testing.T) {

	// declare spread sheet ID and readRange
	spreadsheetId := "12CQBx2eXDX02IyuvLQ8vEE0gSfl_xS24m3vwCKavyCY"
	readRange := "sheet1"

	// test readSheet method
	sheet, err := readSheet(spreadsheetId, readRange)

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(sheet.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("")
		for _, row := range sheet.Values {
			// Print columns A and E, which correspond to indices 0 and 4.
			fmt.Printf(" %s, %s, %s, %s\n", row[0], row[1], row[2], row[3])
		}
	}
}

func TestUpdateSheet(t *testing.T) {
	var err error
	spreadsheetId := "12CQBx2eXDX02IyuvLQ8vEE0gSfl_xS24m3vwCKavyCY"
	var vr sheets.ValueRange
	myval := []interface{}{"test", "http://localhost", "time"}
	vr.Values = append(vr.Values, myval)

	createRange := "sheet1!A2:E"

	// update spread sheet data
	err = updateSheet(spreadsheetId, createRange, vr)
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
}

func TestAddSheet(t *testing.T) {
	var err error
	spreadsheetId := "12CQBx2eXDX02IyuvLQ8vEE0gSfl_xS24m3vwCKavyCY"
	var vr sheets.ValueRange
	myval := []interface{}{"test", "http://localhost", "time"}
	vr.Values = append(vr.Values, myval)

	createRange := "sheet1!A2:E"

	err = addSheet(spreadsheetId, createRange, vr)
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
}
