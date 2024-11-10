package config

import (
	"fmt"
	"os"
	"encoding/csv"
)

// Config is a struct that represents the configuration for the scraper
type Config struct {
	HotelID     string
	Name        string
	City		string
	LocationURL string
	Languages   []string
	FileType    string
	ProxyHost   string
}

func LoadConfig(fileName string) ([]Config, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}
	defer file.Close()
	// Create a new CSV reader
	reader := csv.NewReader(file)
	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Error reading CSV %v", err)
	}
	// Create a slice to hold the extracted records
	var parsedRecords []Config

	// Skip the header row by iterating from index 1
	lang := []string{"en", "it", "de", "es"}
	for _, row := range records {
		record := Config{
			HotelID:     row[0],
			Name:        row[1],
			City:		 row[2],
			LocationURL: row[3],
			Languages:   lang,
			FileType:    "csv",
			ProxyHost:   "",
		}
		parsedRecords = append(parsedRecords, record)
	}
	return parsedRecords, nil
}
