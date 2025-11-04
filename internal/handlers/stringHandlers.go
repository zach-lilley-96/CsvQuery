package handlers

import "strings"

func StrEqualTo(a string, b string) bool {
	return a == b
}

func StrNotEqualTo(a string, b string) bool {
	return a != b
}

func StrContains(a string, b string) bool {
	return strings.Contains(a, b)
}

func StrNotContains(a string, b string) bool {
	return !strings.Contains(a, b)
}
