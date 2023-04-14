package csv

import (
	"CallDigesto/models"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// Write writes a CSV file with the given file name and folder name, and the data from the responses.
func Write(fileName string, folderName string, responses []models.ResponseBody) error {
	// Create a slice to hold all the rows for the CSV file
	var rows [][]string

	// Add the headers to the slice
	rows = append(rows, generateHeaders())

	// Add the data rows to the slice
	for _, response := range responses {
		rows = append(rows, generateRow(response)...)
	}

	// Create the CSV file
	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		log.Println(err)
		return err
	}

	// Close the file when the function completes
	defer cf.Close()

	// Create a new CSV writer
	w := csv.NewWriter(cf)

	// Write all the rows to the CSV file
	err = w.WriteAll(rows)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//createFile function takes in a file path and creates a file in the specified directory. It returns a pointer to the created file and an error if there is any.
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		log.Println(err)
		return nil, err
	}
	return os.Create(p)
}

//generateHeaders function returns a slice of strings containing the header values for the CSV file.
func generateHeaders() []string {
	return []string{
		"Document Number",
		"Document Type",
		"Name",
		"Cover Name",
		"PassivePole",
		"Role",
		"Law Found",
		"Law",
		"Confidence",
		"Lawsuit Number",
		"Main Lawsuit",
		"Link",
		"Lawsuit Year",
		"Distribution Year",
		"Forum",
		"Court City",
		"Court",
		"UF",
		"Justice Type",
		"Nature",
		"Subject",
		"Most Recent Move",
		"Most Recent Update",
	}
}

//generateRow function takes in a single models.WriteStruct argument and returns a slice of strings containing the values to be written in a row of the CSV file.
//	It uses a loop to concatenate all the NameVariations into a single string separated by " | "
func generateRow(response models.ResponseBody) [][]string {
	var rows [][]string

	if len(response.Lawsuits) == 0 {
		// Append Identification
		row := []string{
			response.Identification.Value,
		}
		rows = append(rows, row)
	} else {
		for _, lawsuit := range response.Lawsuits {
			row := []string{
				response.Identification.Value,
				response.Identification.IdType,
				response.Name,
				lawsuit.CoverName,
				strconv.FormatBool(lawsuit.PassivePole),
				lawsuit.Role,
				strconv.FormatBool(lawsuit.CrimeFound),
			}

			// Append specific crime
			if len(lawsuit.Laws) != 0 {
				row = append(row, lawsuit.Laws[0].Crime)
			} else {
				row = append(row, "")
			}

			row = append(row, lawsuit.Confidence)
			row = append(row, lawsuit.LawsuitNumber)
			row = append(row, strconv.FormatBool(lawsuit.MainLawsuit))
			row = append(row, lawsuit.Link)
			row = append(row, lawsuit.LawsuitYear)
			row = append(row, lawsuit.DistYear)
			row = append(row, lawsuit.Forum)
			row = append(row, lawsuit.CourtCity)
			row = append(row, lawsuit.Court)
			row = append(row, lawsuit.UF)
			row = append(row, lawsuit.JusticeType)
			row = append(row, lawsuit.Nature)
			row = append(row, lawsuit.Subject)
			row = append(row, lawsuit.MostRecentMove)
			row = append(row, lawsuit.MostRecentUpdate)

			rows = append(rows, row)
		}
	}
	return rows
}
