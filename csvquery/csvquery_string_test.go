package csvquery

import "testing"

func TestString(t *testing.T) {
	path := "../testFiles/test.csv"

	query := NewCsvQuery(path).SelectAll().Where("teststr").StrEq("hello")
	for _, row := range *query {
		t.Log(row)
	}

	expectedLength := 1
	if len(*query) != expectedLength {
		t.Errorf("Expected %d results, got %d", expectedLength, len(*query))
	}
}
