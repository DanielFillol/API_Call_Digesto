package main

import (
	"CallDigesto/csv"
	"CallDigesto/request"
	"github.com/joho/godotenv"
	"io"
	"log"
	"os"
	"time"
)

const (
	FILENAME      = "response"
	FILEPATH      = "data/requests.csv"
	FILESeparator = ','
	SKIPHeader    = true
	BATCHSize     = 10000
)

func main() {
	// Set log file
	logFile, err := os.Create("output.log.txt")
	if err != nil {
		log.Println("Failed to open log file: " + err.Error())
	}
	defer logFile.Close()

	// Create a multi-writer that writes to both the file and os.Stdout (terminal)
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	// Load environment variables from .env file
	err = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file: ", err)
	}
	// Get the value of the AUTH variable from the environment
	var auth = os.Getenv("AUTH2")

	// Load data to be requested from CSV file
	requests, err := csv.Read(FILEPATH, FILESeparator, SKIPHeader)
	if err != nil {
		log.Println("Error loading requests from CSV: ", err)
	}

	// Make API requests asynchronously
	start := time.Now()
	log.Println("Starting API calls...")

	err = request.AllBatchAsync(requests, BATCHSize, auth, FILENAME)
	if err != nil {
		log.Println(err)
	}

	log.Println("All API calls completed in " + time.Since(start).String())
}
