package main

import (
	"CallDigesto/csv"
	"CallDigesto/request"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

const LINK = "https://op.digesto.com.br/api/background_check/advanced_search_all?api_key="
const METHOD = "POST"
const WORKERS = 1

func main() {
	//load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env file was not found. You should add a .env file on project root with:\nAUTH")
	}

	//get .env variables
	var auth = os.Getenv("AUTH")

	//load data to be requested
	requests, err := csv.Read("data/requests.csv", ',')
	if err != nil {
		fmt.Println(err)
	}

	//requests
	url := LINK + auth
	start := time.Now()
	results, err := request.AsyncAPIRequest(requests, WORKERS, url, METHOD, auth)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("finished in ", time.Since(start))

	//download API response to files
	err = csv.Write("result", "data", results)
	if err != nil {
		fmt.Println(err)
	}

}
