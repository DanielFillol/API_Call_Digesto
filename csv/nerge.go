package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const MergedFilename = "merged_result.csv"

func MergeAndDeleteCSVs(folderPath string) error {
	// List all CSV files in the specified folder
	files, err := filepath.Glob(filepath.Join(folderPath, "*.csv"))
	if err != nil {
		return err
	}

	if len(files) == 0 {
		log.Println("No CSV files found in the specified folder.")
		return nil
	}

	// Create or open the merged CSV file
	mergedFile, err := os.Create(MergedFilename)
	if err != nil {
		return err
	}
	defer mergedFile.Close()

	// Create a CSV writer for the merged file
	mergedWriter := csv.NewWriter(mergedFile)
	defer mergedWriter.Flush()

	// Write headers from the first CSV file
	firstFile, err := os.Open(files[0])
	if err != nil {
		return err
	}
	defer firstFile.Close()

	firstReader := csv.NewReader(firstFile)
	headers, err := firstReader.Read()
	if err != nil {
		return err
	}
	mergedWriter.Write(headers)

	// Iterate through each CSV file, skipping the header in subsequent files
	for _, file := range files {
		if file == MergedFilename {
			continue // Skip the merged file itself
		}

		currentFile, err := os.Open(file)
		if err != nil {
			return err
		}
		defer currentFile.Close()

		currentReader := csv.NewReader(currentFile)
		if _, err := currentReader.Read(); err != nil && err != io.EOF {
			return err
		}

		// Write the remaining rows to the merged file
		lineNumber := 2 // Start from line 2 (skipping header)
		for {
			row, err := currentReader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(row)
				//return err
			}

			// Check if the number of fields in the row matches the number of headers
			if len(row) != len(headers) {
				// If the number of fields is less than headers, pad with empty strings
				if len(row) < len(headers) {
					for i := len(row); i < len(headers); i++ {
						row = append(row, "")
					}
				} else {
					// If the number of fields is more than headers, truncate the record
					row = row[:len(headers)]
				}
			}

			if err := mergedWriter.Write(row); err != nil {
				log.Println(err)
				return err
			}

			lineNumber++
		}
	}

	// Remove individual CSV files
	for _, file := range files {
		if file == MergedFilename {
			continue // Skip the merged file itself
		}
		if err := os.Remove(file); err != nil {
			log.Printf("Error deleting file %s: %v", file, err)
		}
	}

	log.Printf("Merged CSV file created: %s", MergedFilename)

	return nil
}
