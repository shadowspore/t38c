package t38c

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type cmd struct {
	Name string
	Args []string
}

func newCmd(name string, args ...string) cmd {
	return cmd{name, args}
}

func (c cmd) String() string {
	str := c.Name
	if len(c.Args) > 0 {
		str += " " + strings.Join(c.Args, " ")
	}
	return str
}

func floatString(val float64) string {
	return strconv.FormatFloat(val, 'f', -1, 64)
}

func rawEventHandler(handler func(*GeofenceEvent) error) func([]byte) error {
	return func(data []byte) error {
		resp := &GeofenceEvent{}
		if err := json.Unmarshal(data, resp); err != nil {
			return fmt.Errorf("json unmarshal geofence response: %v", err)
		}

		return handler(resp)
	}
}
