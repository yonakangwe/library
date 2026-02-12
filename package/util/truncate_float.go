package util

import "math"

// TruncateFloat64
func TruncateFloat64(value float64, precision int) float64 {
	return math.Floor(value*math.Pow10(precision)) / math.Pow10(precision)
}
