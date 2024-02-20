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
	APIBNMP = "https://api-dev.consulta-pro.jusbrasil.com.br/other-records"
	API     = "/api/background_check/advanced_search_all"
	BASE    = "https://op.digesto.com.br"
	METHOD  = "POST"
	WORKERS = 1
)

const (
	FILEPATH      = "data/requests.csv"
	FILESEPARATOR = ','
	SKIPHEADER    = true
)

const (
	FILENAME = "response_Criminal_and_Civil"
	FOLDER   = "data"
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

	results, err := request.AsyncAPIRequest(requests, WORKERS, urlCaller, METHOD, auth)
	if err != nil {
		log.Println("Error making API requests: ", err)
	}

	log.Println("Finished API_Criminal calls in ", time.Since(start))

	// Write API response to CSV file
	err = csv.Write(FILENAME, FOLDER, results)
	if err != nil {
		log.Fatal("Error writing API response to CSV: ", err)
	}

	//Starts Other endpoint request
	urlCaller = APIBNMP
	auth = os.Getenv("AUTH2")

	results2, err := request.AsyncAPIRequestBNMP(requests, WORKERS, urlCaller, METHOD, auth)
	if err != nil {
		log.Println("Error making API requests: ", err)
	}

	// Write API response to CSV file
	err = csv.WriteOthers(FOLDER, results2)
	if err != nil {
		log.Fatal("Error writing API response to CSV: ", err)
	}

	log.Println("Finished API_Others calls in ", time.Since(start))

}
