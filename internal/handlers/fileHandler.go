package handlers

import (
	"encoding/csv"
	"os"
)

func GetHeaders(path string) ([]string, uint) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.Read()
	if err != nil {
		panic(err)
	}

	return data, uint(len(data))
}

func GetAll(path string) *[][]string {
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

	return &data
}
