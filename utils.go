package t38c

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func floatString(val float64) string {
	return strconv.FormatFloat(val, 'g', 10, 64)
}

func rawEventHandler(handler func(*GeofenceEvent)) func([]byte) error {
	return func(data []byte) error {
		resp := &GeofenceEvent{}
		if err := json.Unmarshal(data, resp); err != nil {
			return fmt.Errorf("json unmarshal geofence response: %v", err)
		}

		handler(resp)
		return nil
	}
}
