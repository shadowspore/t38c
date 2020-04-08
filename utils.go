package t38c

import (
	"strconv"
)

func floatString(val float64) string {
	return strconv.FormatFloat(val, 'f', 10, 64)
}
