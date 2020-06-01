package t38c

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type tileCmd struct {
	Name string
	Args []string
}

func newTileCmd(name string, args ...string) tileCmd {
	return tileCmd{name, args}
}

func (cmd tileCmd) appendArgs(name string, args ...string) tileCmd {
	cmd.Args = append(cmd.Args, args...)
	return cmd
}

func (cmd tileCmd) String() string {
	str := cmd.Name
	for _, arg := range cmd.Args {
		str += " " + arg
	}

	return str
}

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
