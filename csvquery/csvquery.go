package csvquery

import (
	"fmt"
	"slices"

	handler "github.com/zach-lilley-96/csvquery/internal"
)

type CsvQuery struct {
	Path            string
	headers         []string
	numberOfRecords uint
	selectedColum   string
	getAll          bool
	strCompareTo    *string
	numCompareTo    *float64
}

func NewCsvQuery(path string) *CsvQuery {
	headers := handler.GetHeaders(path)
	return &CsvQuery{
		Path:         path,
		headers:      headers,
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
	return cq
}

func (cq *CsvQuery) Where(column string) *CsvQuery {
	if !slices.Contains(cq.headers, column) {
		error := fmt.Sprintf("Column %s does not exist", column)
		panic(error)
	}
	cq.selectedColum = column
	return cq
}

func (cq *CsvQuery) Gt(value float64) {
	cq.numCompareTo = &value
}

func (cq *CsvQuery) Gte(value float64) {
	cq.numCompareTo = &value
}

func (cq *CsvQuery) Lt(value float64) {
	cq.numCompareTo = &value
}

func (cq *CsvQuery) Lte(value float64) {
	cq.numCompareTo = &value
}

func (cq *CsvQuery) Eq(value float64) {
	cq.numCompareTo = &value
}

func (cq *CsvQuery) StrEq(value string) {
	cq.strCompareTo = &value
}

func (cq *CsvQuery) Ct(value string) {
	cq.strCompareTo = &value
}

func (cq *CsvQuery) GetAll() {

}

func (cq *CsvQuery) Get(quantity int) {

}

func (cq *CsvQuery) First() {

}
