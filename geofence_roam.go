package t38c

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func buildRoamArgs(req *GeofenceRequest) []string {
	var args []string
	args = append(args, req.key)

	for _, opt := range req.opts {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}

	args = append(args, "FENCE")

	if len(req.actions) > 0 {
		args = append(args, "DETECT")
		actions := ""
		first := true
		for _, action := range req.actions {
			if !first {
				actions += ","
			}
			actions += string(action)
			first = false
		}
		args = append(args, actions)
	}

	if len(req.objectType) > 0 {
		args = append(args, req.objectType)
	}

	args = append(args, []string{
		"ROAM", req.target, req.pattern, strconv.Itoa(req.meters),
	}...)

	return args
}

// FenceRoam ...
func (client *FenceClient) FenceRoam(req *GeofenceRequest) (GeofenceRoamObjectChan, error) {
	args := buildRoamArgs(req)
	ch, err := client.FenceExecute("NEARBY", args...)
	if err != nil {
		return nil, err
	}

	objChan := make(GeofenceRoamObjectChan, 10)
	go func() {
		defer close(objChan)
		for event := range ch {
			if event.Err != nil {
				log.Fatal(event.Err)
			}
			fmt.Printf("%s\n", event.Data)
			var resp GeofenceRoamObject
			if err := json.Unmarshal(event.Data, &resp); err != nil {
				panic(err)
			}

			objChan <- resp
		}
	}()

	return objChan, nil
}
