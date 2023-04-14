package main

import (
	"CallDigesto/csv"
	"CallDigesto/request"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

const (
	LINK    = "https://op.digesto.com.br/api/background_check/advanced_search_all?api_key="
	METHOD  = "POST"
	WORKERS = 2
)

const (
	FILEPATH      = "data/requests.csv"
	FILESEPARATOR = ','
	SKIPHEADER    = true
)

const (
	FILENAME = "response"
	FOLDER   = "data"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	// Get the value of the AUTH variable from the environment
	var auth = os.Getenv("AUTH")

	// Load data to be requested from CSV file
	requests, err := csv.Read(FILEPATH, FILESEPARATOR, SKIPHEADER)
	if err != nil {
		log.Fatal("Error loading requests from CSV: ", err)
	}

	// Make API requests asynchronously
	start := time.Now()
	log.Println("Starting API calls...")
	url := LINK + auth
	results, err := request.AsyncAPIRequest(requests, WORKERS, url, METHOD, auth)
	if err != nil {
		log.Fatal("Error making API requests: ", err)
	}
	log.Println("Finished API calls in ", time.Since(start))

	// Write API response to CSV file
	err = csv.Write(FILENAME, FOLDER, results)
	if err != nil {
		log.Fatal("Error writing API response to CSV: ", err)
	}
}
