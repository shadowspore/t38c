package tests

import (
	"fmt"
	"strconv"
	"strings"
)

func requestToString(cmd string, args []interface{}) string {
	var sb strings.Builder
	sb.WriteString(cmd)
	for _, arg := range args {
		sb.WriteString(" " + ifaceToString(arg))
	}

	return sb.String()
}

func ifaceToString(i interface{}) string {
	switch i := i.(type) {
	case int:
		return strconv.Itoa(i)
	case float64:
		return strconv.FormatFloat(i, 'f', 10, 64)
	case string:
		return i
	default:
		return fmt.Sprintf("%s", i)
	}
}
