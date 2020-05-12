package t38c

import (
	"fmt"
	"strconv"

	"github.com/tidwall/gjson"
)

func floatString(val float64) string {
	return strconv.FormatFloat(val, 'f', 10, 64)
}

func checkResponseErr(resp []byte) error {
	if !gjson.GetBytes(resp, "ok").Bool() {
		return fmt.Errorf(gjson.GetBytes(resp, "err").String())
	}

	return nil
}
