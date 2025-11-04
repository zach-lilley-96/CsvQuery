package csvquery

import (
	"fmt"
	"testing"
)

func TestHandler(t *testing.T) {
	path := "../testFiles/test.csv"

	query := NewCsvQuery(path).SelectAll().Where("testint").Gte(2.17)
	for _, row := range *query {
		fmt.Println(row)
	}

	expectedLength := 9
	if len(*query) != expectedLength {
		t.Errorf("Expected %d results, got %d", expectedLength, len(*query))
	}

	query2 := NewCsvQuery(path).SelectAll().Where("testint").Lt(2.17)
	for _, row := range *query2 {
		fmt.Println(row)
	}

	expectedLength2 := 5
	if len(*query2) != expectedLength2 {
		t.Errorf("Expected %d results, got %d", expectedLength2, len(*query2))
	}
}
