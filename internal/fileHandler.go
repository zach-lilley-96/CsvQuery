package internal

import (
	"encoding/csv"
	"os"
)

func GetCsvData(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return data

}
