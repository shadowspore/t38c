package t38c

import (
	"strconv"
)

func floatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', 6, 64)
}
