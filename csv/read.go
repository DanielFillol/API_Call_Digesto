package csv

import (
	"CallDigesto/models"
	"encoding/csv"
	"os"
)

func Read(filePath string, separator rune) ([]models.ReadCsv, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvR := csv.NewReader(csvFile)
	csvR.Comma = separator

	csvData, err := csvR.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []models.ReadCsv
	for i, line := range csvData {
		if i != 0 {
			document := line[0]
			data = append(data, models.ReadCsv{
				Document: document,
			})
		}
	}

	return data, nil
}
