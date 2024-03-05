package main

import (
	"CallDigesto/csv"
	"CallDigesto/request"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	API           = "/api/background_check/advanced_search_all"
	BASE          = "https://op.digesto.com.br"
	METHOD        = "POST"
	WORKERS       = 1
	BATCHSize     = 10000
	BATCHInterval = 5 * time.Second
)

const (
	FILEPATH      = "data/requests.csv"
	FILESEPARATOR = ','
	SKIPHEADER    = true
	FILENAME      = "response_Criminal_and_Civil"
	FOLDER        = "data/response"
)

func main() {
	var urlCaller = BASE + API + "?api_key="

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

	// Process requests in batches
	var resultsSaved int
	for i := 0; i < len(requests); i += BATCHSize {
		end := i + BATCHSize
		if end > len(requests) {
			end = len(requests)
		}

		batchRequests := requests[i:end]

		// Make API requests asynchronously
		start := time.Now()
		log.Printf("Starting API calls for batch %d...", i/BATCHSize)

		batchResults, err := request.AsyncAPIRequest(batchRequests, WORKERS, urlCaller, METHOD, auth)
		if err != nil {
			log.Printf("Error making API requests for batch %d: %v", i/BATCHSize, err)
			continue
		}

		log.Printf("Finished API calls for batch %d in %v", i/BATCHSize, time.Since(start))

		// Write API response to CSV file
		err = csv.Write(FILENAME+"_"+strconv.Itoa(i/BATCHSize), FOLDER, batchResults)
		if err != nil {
			log.Printf("Error writing API response to CSV for batch %d: %v", i/BATCHSize, err)
			continue
		}

		resultsSaved += len(batchResults)
		// Introduce a 5-second interval between batches
		time.Sleep(BATCHInterval)
	}

	log.Println("All API calls completed in " + time.Since(start).String())

	err = csv.MergeAndDeleteCSVs(FOLDER)
	if err != nil {
		log.Println("Error merging csv: ", err)
	}
}
