package csvquery

import (
	"fmt"
	"slices"
	"strconv"

	handler "github.com/zach-lilley-96/csvquery/internal/handlers"
)

type CsvQuery struct {
	Path            string
	headers         []string
	length          uint
	numberOfRecords uint
	selectedColumn  int
	getAll          bool
	strCompareTo    *string
	numCompareTo    *float64
}

func NewCsvQuery(path string) *CsvQuery {
	headers, len := handler.GetHeaders(path)
	return &CsvQuery{
		Path:         path,
		headers:      headers,
		length:       len,
		getAll:       false,
		strCompareTo: nil,
		numCompareTo: nil,
	}
}

func (cq *CsvQuery) Select(quantity uint) *CsvQuery {
	cq.numberOfRecords = quantity
	return cq
}

func (cq *CsvQuery) SelectAll() *CsvQuery {
	cq.getAll = true
	cq.numberOfRecords = cq.length
	return cq
}

func (cq *CsvQuery) Where(column string) *CsvQuery {
	if !slices.Contains(cq.headers, column) {
		error := fmt.Sprintf("Column %s does not exist", column)
		panic(error)
	}
	cq.selectedColumn = slices.Index(cq.headers, column)
	return cq
}

func (cq *CsvQuery) numberQuery(fn func(float64, float64) bool) *[][]string {
	var results [][]string
	fileData := handler.GetAll(cq.Path)
	count := 0
	for i, row := range *fileData {
		if count >= int(cq.numberOfRecords) && !cq.getAll {
			break
		}
		if i != 0 {
			valueAsFloat, err := strconv.ParseFloat(row[cq.selectedColumn], 64)
			if err != nil {
				continue
			}

			if fn(valueAsFloat, *cq.numCompareTo) {
				results = append(results, row)
				count++
			}
		}
	}
	return &results
}

func (cq *CsvQuery) strQuery(fn func(string, string) bool) *[][]string {
	if cq.strCompareTo == nil {
		panic("String to compare to is nil")
	}
	var results [][]string
	fileData := handler.GetAll(cq.Path)
	for i, row := range *fileData {
		if i != 0 {

			if fn(row[cq.selectedColumn], *cq.strCompareTo) {
				results = append(results, row)
			}
		}
	}
	return &results
}

func (cq *CsvQuery) Gt(value float64) *[][]string {
	cq.numCompareTo = &value
	return cq.numberQuery(handler.GreaterThan)
}

func (cq *CsvQuery) Gte(value float64) *[][]string {
	cq.numCompareTo = &value
	return cq.numberQuery(handler.GreaterThanOrEqual)
}

func (cq *CsvQuery) Lt(value float64) *[][]string {
	cq.numCompareTo = &value
	return cq.numberQuery(handler.LessThan)
}

func (cq *CsvQuery) Lte(value float64) *[][]string {
	cq.numCompareTo = &value
	return cq.numberQuery(handler.LessThanOrEqual)
}

func (cq *CsvQuery) Eq(value float64) *[][]string {
	cq.numCompareTo = &value
	return cq.numberQuery(handler.EqualTo)
}

func (cq *CsvQuery) StrEq(value string) *[][]string {
	cq.strCompareTo = &value
	return cq.strQuery(handler.StrEqualTo)
}

func (cq *CsvQuery) StrNeq(value string) *[][]string {
	cq.strCompareTo = &value
	return cq.strQuery(handler.StrNotEqualTo)
}

func (cq *CsvQuery) StrContains(value string) *[][]string {
	cq.strCompareTo = &value
	return cq.strQuery(handler.StrContains)
}

func (cq *CsvQuery) StrNotContains(value string) *[][]string {
	cq.strCompareTo = &value
	return cq.strQuery(handler.StrNotContains)
}
