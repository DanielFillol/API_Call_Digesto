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
	API     = "/lawsuits/all"
	BASE    = "https://api.consulta-pro.jusbrasil.com.br"
	METHOD  = "POST"
	WORKERS = 1
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
	log.Println("Finished API calls in ", time.Since(start))

	// Write API response to CSV file
	err = csv.Write(FILENAME, FOLDER, results)
	if err != nil {
		log.Fatal("Error writing API response to CSV: ", err)
	}
}
