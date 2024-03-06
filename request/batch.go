package request

import (
	"CallDigesto/csv"
	"CallDigesto/models"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	METHOD = "POST"
)
const (
	criminalCall           = "/lawsuits/criminal"
	criminalFolderSingle   = "data/response/criminal"
	MergedFilenameCriminal = "merged_result_criminal.csv"
)
const (
	civilCall           = "/lawsuits/civil"
	civilFolderSingle   = "data/response/civil"
	MergedFilenameCivil = "merged_result_civil.csv"
)
const (
	criminalOtherCall           = "/other-records/criminal"
	criminalOtherFolderSingle   = "data/response/other_records/criminal"
	MergedFilenameOtherCriminal = "merged_result_criminal.csv"
)
const (
	civilOtherCall           = "/other-records/civel"
	civilOtherFolderSingle   = "data/response/other_records/civil"
	MergedFilenameOtherCivil = "merged_result_other_civil.csv"
)

var WORKERS int
var BATCHInterval time.Duration
var CollDown time.Duration

func AllBatchAsync(requests []models.ReadCsv, batchSize int, auth string, fileName string) error {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		err := batchCall(requests, batchSize, criminalCall, auth, criminalFolderSingle, criminalFolderSingle, fileName, MergedFilenameCriminal)
		if err != nil {
			log.Print("Error on criminal caller: ", err)
		}
	}()

	go func() {
		//	defer wg.Done()
		err := batchCall(requests, batchSize, civilCall, auth, civilFolderSingle, civilFolderSingle, fileName, MergedFilenameCivil)
		if err != nil {
			log.Print("Error on civil caller: ", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := batchCallOthers(requests, batchSize, criminalOtherCall, auth, criminalOtherFolderSingle, criminalOtherFolderSingle, fileName, MergedFilenameOtherCriminal)
		if err != nil {
			log.Print("Error on criminal other caller: ", err)
		}
	}()

	go func() {
		defer wg.Done()
		err := batchCallOthers(requests, batchSize, civilOtherCall, auth, civilOtherFolderSingle, civilOtherFolderSingle, fileName, MergedFilenameOtherCivil)
		if err != nil {
			log.Print("Error on civil other caller: ", err)
		}
	}()

	wg.Wait()

	return nil
}

func batchCall(requests []models.ReadCsv, batchSize int, call string, auth string, folderMerge string, folderName string, fileName string, mergedFileName string) error {
	var urlCaller = "https://api-dev.consulta-pro.jusbrasil.com.br" + call

	// Process requests in batches
	var resultsSaved int
	for i := 0; i < len(requests); i += batchSize {
		end := i + batchSize
		if end > len(requests) {
			end = len(requests)
		}

		batchRequests := requests[i:end]

		// Make API requests asynchronously
		start := time.Now()
		log.Printf("Starting API calls for "+strings.ReplaceAll(call, "/", " ")+" batch %d...", i/batchSize)

		batchResults, err := AsyncAPIRequest(batchRequests, WORKERS, urlCaller, METHOD, auth, CollDown)
		if err != nil {
			log.Printf("Error making API "+strings.ReplaceAll(call, "/", " ")+" requests for batch %d: %v", i/batchSize, err)
			continue
		}

		log.Printf("Finished API calls for "+strings.ReplaceAll(call, "/", " ")+" batch %d in %v", i/batchSize, time.Since(start))

		// WriteLawsuits API response to CSV file
		err = csv.WriteLawsuits(fileName+"_"+strconv.Itoa(i/batchSize), folderName, batchResults)
		if err != nil {
			log.Printf("Error writing API "+strings.ReplaceAll(call, "/", " ")+" response to CSV for batch %d: %v", i/batchSize, err)
			continue
		}

		resultsSaved += len(batchResults)
		// Introduce interval between batches
		time.Sleep(BATCHInterval)
	}

	if resultsSaved != 0 {
		err := csv.MergeAndDeleteCSVs(folderMerge, mergedFileName)
		if err != nil {
			return err
		}
	}

	return nil
}

func batchCallOthers(requests []models.ReadCsv, BATCHSize int, call string, auth string, folderMerge string, folderName string, fileName string, mergedFileName string) error {
	var urlCaller = "https://api-dev.consulta-pro.jusbrasil.com.br" + call

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
		log.Printf("Starting API calls for "+strings.ReplaceAll(call, "/", " ")+" batch %d...", i/BATCHSize)

		batchResults, err := AsyncAPIRequestOther(batchRequests, WORKERS, urlCaller, METHOD, auth, CollDown)
		if err != nil {
			log.Printf("Error making API "+strings.ReplaceAll(call, "/", " ")+" requests for batch %d: %v", i/BATCHSize, err)
			continue
		}

		log.Printf("Finished API calls for "+strings.ReplaceAll(call, "/", " ")+" batch %d in %v", i/BATCHSize, time.Since(start))

		// WriteLawsuits API response to CSV file
		err = csv.WriteOthers("_"+strconv.Itoa(i/BATCHSize), folderName, batchResults)
		if err != nil {
			log.Printf("Error writing API "+strings.ReplaceAll(call, "/", " ")+" response to CSV for batch %d: %v", i/BATCHSize, err)
			continue
		}

		resultsSaved += len(batchResults)
		// Introduce interval between batches
		time.Sleep(BATCHInterval)
	}

	if resultsSaved != 0 {
		err := csv.MergeAndDeleteCSVs(folderMerge, mergedFileName)
		if err != nil {
			return err
		}
	}

	return nil
}
