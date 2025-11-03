package handlers

import "strconv"

func GreaterThan(path string, columnIndex int, compValue float64, quantity int) *[][]string {
	var results [][]string
	fileData := GetAll(path)
	for i, row := range *fileData {
		if i != 0 {
			valueAsFloat, err := strconv.ParseFloat(row[columnIndex], 64)
			if err != nil {
				continue
			}

			if valueAsFloat > compValue {
				results = append(results, row)
			}
		}
	}
	return &results
}

func GreaterThanOrEqual(path string, columnIndex int, compValue float64, quantity int) *[][]string {
	var results [][]string
	fileData := GetAll(path)
	for i, row := range *fileData {
		if i != 0 {
			valueAsFloat, err := strconv.ParseFloat(row[columnIndex], 64)
			if err != nil {
				continue
			}

			if valueAsFloat >= compValue {
				results = append(results, row)
			}
		}
	}
	return &results
}

func LessThan(path string, columnIndex int, compValue float64, quantity int) *[][]string {
	var results [][]string
	fileData := GetAll(path)
	for i, row := range *fileData {
		if i != 0 {
			valueAsFloat, err := strconv.ParseFloat(row[columnIndex], 64)
			if err != nil {
				continue
			}

			if valueAsFloat < compValue {
				results = append(results, row)
			}
		}
	}
	return &results
}

func LessThanOrEqual(path string, columnIndex int, compValue float64, quantity int) *[][]string {
	var results [][]string
	fileData := GetAll(path)
	for i, row := range *fileData {
		if i != 0 {
			valueAsFloat, err := strconv.ParseFloat(row[columnIndex], 64)
			if err != nil {
				continue
			}

			if valueAsFloat <= compValue {
				results = append(results, row)
			}
		}
	}
	return &results
}

func EqualTo(path string, columnIndex int, compValue float64, quantity int) *[][]string {
	var results [][]string
	fileData := GetAll(path)
	for i, row := range *fileData {
		if i != 0 {
			valueAsFloat, err := strconv.ParseFloat(row[columnIndex], 64)
			if err != nil {
				continue
			}

			if valueAsFloat == compValue {
				results = append(results, row)
			}
		}
	}
	return &results
}
