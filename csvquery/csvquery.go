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

func (cq *CsvQuery) Execute(fn func(float64, float64) bool) *[][]string {
	var results [][]string
	fileData := handler.GetAll(cq.Path)
	for i, row := range *fileData {
		if i != 0 {
			valueAsFloat, err := strconv.ParseFloat(row[cq.selectedColumn], 64)
			if err != nil {
				continue
			}

			if fn(valueAsFloat, *cq.numCompareTo) {
				results = append(results, row)
			}
		}
	}
	return &results
}

func (cq *CsvQuery) Gt(value float64) *[][]string {
	cq.numCompareTo = &value
	return cq.Execute(handler.GreaterThan)
}

func (cq *CsvQuery) Gte(value float64) *[][]string {
	cq.numCompareTo = &value
	return cq.Execute(handler.GreaterThanOrEqual)
}

func (cq *CsvQuery) Lt(value float64) *[][]string {
	cq.numCompareTo = &value
	return cq.Execute(handler.LessThan)
}

func (cq *CsvQuery) Lte(value float64) *[][]string {
	cq.numCompareTo = &value
	return cq.Execute(handler.LessThanOrEqual)
}

func (cq *CsvQuery) Eq(value float64) *[][]string {
	cq.numCompareTo = &value
	return cq.Execute(handler.EqualTo)
}

func (cq *CsvQuery) StrEq(value string) {
	cq.strCompareTo = &value
}

func (cq *CsvQuery) Ct(value string) {
	cq.strCompareTo = &value
}
