package csvquery

import (
	"testing"
)

func TestHandler(t *testing.T) {
	path := "../testFiles/test.csv"

	query := NewCsvQuery(path).SelectAll().Where("testint").Gte(2.17)
	for _, row := range *query {
		t.Log(row)
	}

	expectedLength := 3
	if len(*query) != expectedLength {
		t.Errorf("Expected %d results, got %d", expectedLength, len(*query))
	}
}
