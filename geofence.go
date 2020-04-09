package t38c

import (
	"encoding/json"
)

// FenceObject ...
func FenceObject(exec GeofenceExecutor, req GeofenceRequestable) (GeofenceRoamObjectChan, error) {
	cmd := req.GeofenceCommand()
	ch, err := exec.Fence(cmd.Name, cmd.Args...)
	if err != nil {
		return nil, err
	}

	objChan := make(GeofenceRoamObjectChan, 10)
	go func() {
		defer close(objChan)
		for event := range ch {
			var resp GeofenceRoamObject
			if err := json.Unmarshal(event, &resp); err != nil {
				panic(err)
			}

			objChan <- resp
		}
	}()

	return objChan, nil
}
