package t38c

import (
	"encoding/json"
	"fmt"
	"log"
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

func unmarshalEvents(events chan []byte) (chan GeofenceResponse, error) {
	ch := make(chan GeofenceResponse, 10)
	go func() {
		defer close(ch)
		for event := range events {
			var resp GeofenceResponse
			if err := json.Unmarshal(event, &resp); err != nil {
				log.Printf("bad event: %v", err)
				break
			}

			ch <- resp
		}
	}()

	return ch, nil
}
