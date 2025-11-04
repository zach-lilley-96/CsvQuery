package handlers

func GreaterThan(value float64, comp float64) bool {
	return value > comp
}

func GreaterThanOrEqual(value float64, comp float64) bool {
	return value >= comp
}

func LessThan(value float64, comp float64) bool {
	return value < comp
}

func LessThanOrEqual(value float64, comp float64) bool {
	return value <= comp
}

func EqualTo(value float64, comp float64) bool {
	return value == comp
}
