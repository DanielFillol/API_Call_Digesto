package csv

import (
	"CallDigesto/models"
	"encoding/csv"
	"log"
	"strconv"
)

func WriteOthers(complementFile string, folderName string, responses []models.ResponseBodyOtherRecords) error {
	err := writeMP(complementFile, folderName+"/mp", responses)
	if err != nil {
		return err
	}
	err = writeBNMP(complementFile, folderName+"/bnmp", responses)
	if err != nil {
		return err
	}

	return nil
}

func writeBNMP(complementFile string, folderName string, responses []models.ResponseBodyOtherRecords) error {
	// Create a slice to hold all the rows for the CSV file
	var rows [][]string

	// Add the headers to the slice
	rows = append(rows, generateHeadersOtherBNMP())

	// Add the data rows to the slice
	for _, response := range responses {
		rows = append(rows, generateRowOtherBNMP(response)...)
	}

	// Create the CSV file
	cf, err := createFile(folderName + "/" + "response_bnmp" + complementFile + ".csv")
	if err != nil {
		log.Println(err)
		return err
	}

	// Close the file when the function completes
	defer cf.Close()

	// Create a new CSV writer
	w := csv.NewWriter(cf)

	// WriteLawsuits all the rows to the CSV file
	err = w.WriteAll(rows)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// generateHeadersOther function returns a slice of strings containing the header values for the CSV file.
func generateHeadersOtherBNMP() []string {
	return []string{
		"Document Number",
		"Document Type",
		"Name",
		"BNMP Sigill",
		"BNMP Confidence",
		"BNMP UF",
		"BNMP Situation",
		"BNMP Organ",
		"BNMP Prison Type",
		"BNMP Document Number",
		"BNMP Document Expedition",
		"BNMP Other Name",
		"BNMP Mother Name",
		"BNMP Father Name",
		"BNMP Birth Date",
		"BNMP Nationality",
		"BNMP Place Of Birth",
		"BNMP Profession",
		"BNMP Lawsuit Location",
		"BNMP Judge",
		"BNMP Validity Date",
		"BNMP Creation Date",
		"BNMP Recapture",
		"BNMP Decision",
		"BNMP Execution",
		"BNMP Observation",
		"BNMP Place Of Crime ",
		"BNMP Penalty Time",
		"BNMP Regime",
		"BNMP Crime Found",
		"BNMP Laws",
	}
}

// generateRow function takes in a single models.WriteStruct argument and returns a slice of strings containing the values to be written in a row of the CSV file.
func generateRowOtherBNMP(response models.ResponseBodyOtherRecords) [][]string {
	var rows [][]string

	if len(response.BNMP) == 0 {
		// Append Identification
		row := []string{
			response.Identification.Value,
		}
		rows = append(rows, row)
	} else {
		for _, lawsuit := range response.BNMP {
			row := []string{
				response.Identification.Value,
				response.Identification.IdType,
				response.Name,
			}

			row = append(row, lawsuit.Sigill)
			row = append(row, lawsuit.Confidence)
			row = append(row, lawsuit.UF)
			row = append(row, lawsuit.Situation)
			row = append(row, lawsuit.Organ)
			row = append(row, lawsuit.PrisonType)
			row = append(row, lawsuit.DocumentNumber)
			row = append(row, lawsuit.DocumentExpedition)
			row = append(row, lawsuit.OtherName)
			row = append(row, lawsuit.MotherName)
			row = append(row, lawsuit.FatherName)
			row = append(row, lawsuit.BirthDate)
			row = append(row, lawsuit.Nationality)
			row = append(row, lawsuit.PlaceOfBirth)
			row = append(row, lawsuit.Profession)
			row = append(row, lawsuit.LawsuitLocation)
			row = append(row, lawsuit.Judge)
			row = append(row, lawsuit.ValidityDate)
			row = append(row, lawsuit.CreationDate)
			row = append(row, strconv.FormatBool(lawsuit.Recapture))
			row = append(row, lawsuit.Decision)
			row = append(row, lawsuit.Execution)
			row = append(row, lawsuit.Observation)
			row = append(row, lawsuit.PlaceOfCrime)
			row = append(row, lawsuit.PenaltyTime)
			row = append(row, lawsuit.Regime)
			row = append(row, strconv.FormatBool(lawsuit.CrimeFound))
			// Append specific crime
			if len(lawsuit.Laws) != 0 {
				row = append(row, lawsuit.Laws[0].Crime)
			} else {
				row = append(row, "")
			}
			rows = append(rows, row)
		}
	}
	return rows
}

func writeMP(complementFile string, folderName string, responses []models.ResponseBodyOtherRecords) error {
	// Create a slice to hold all the rows for the CSV file
	var rows [][]string

	// Add the headers to the slice
	rows = append(rows, generateHeadersOtherMP())

	// Add the data rows to the slice
	for _, response := range responses {
		rows = append(rows, generateRowOtherMP(response)...)
	}

	// Create the CSV file
	cf, err := createFile(folderName + "/" + "response_mp" + complementFile + ".csv")
	if err != nil {
		log.Println(err)
		return err
	}

	// Close the file when the function completes
	defer cf.Close()

	// Create a new CSV writer
	w := csv.NewWriter(cf)

	// WriteLawsuits all the rows to the CSV file
	err = w.WriteAll(rows)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// generateHeadersOther function returns a slice of strings containing the header values for the CSV file.
func generateHeadersOtherMP() []string {
	return []string{
		"Document Number",
		"Document Type",
		"Name",
		"Sigil",
		"Confidence",
		"Subject",
		"MPUnity",
		"NumberMP",
		"LawsuitNumber",
		"ProcedureType",
		"Situation",
		"YearLawsuit",
		"UF",
		"Autos",
		"Law",
		"Law Found",
	}
}

// generateRowOther function takes in a single models.ResponseBodyOtherRecords argument
// and returns a slice of strings containing the values to be written in a row of the CSV file.
func generateRowOtherMP(response models.ResponseBodyOtherRecords) [][]string {
	var rows [][]string

	if len(response.MP) == 0 {
		// Append Identification
		row := []string{
			response.Identification.Value,
		}
		rows = append(rows, row)
	} else {
		for _, lawsuit := range response.MP {
			row := []string{
				response.Identification.Value,
				response.Identification.IdType,
				response.Name,
			}
			row = append(row, lawsuit.Sigill)
			row = append(row, lawsuit.Confidence)
			row = append(row, lawsuit.Subject)
			row = append(row, lawsuit.MPUnity)
			row = append(row, lawsuit.NumberMP)
			row = append(row, lawsuit.LawsuitNumber)
			row = append(row, lawsuit.ProcedureType)
			row = append(row, lawsuit.Situation)
			row = append(row, lawsuit.YearLawsuit)
			row = append(row, lawsuit.UF)
			row = append(row, lawsuit.Autos)
			// Append specific crime
			if len(lawsuit.Laws) != 0 {
				row = append(row, lawsuit.Laws[0].Crime)
			} else {
				row = append(row, "")
			}
			row = append(row, strconv.FormatBool(lawsuit.CrimeFound))
			rows = append(rows, row)
		}
	}
	return rows
}
