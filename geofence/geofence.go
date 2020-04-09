package geofence

import (
	"encoding/json"
)

// FenceObject ...
func FenceObject(exec Executor, req Requestable) (RoamObjectChan, error) {
	cmd := req.GeofenceCommand()
	ch, err := exec.Fence(cmd.Name, cmd.Args...)
	if err != nil {
		return nil, err
	}

	objChan := make(RoamObjectChan, 10)
	go func() {
		defer close(objChan)
		for event := range ch {
			var resp RoamObject
			if err := json.Unmarshal(event, &resp); err != nil {
				panic(err)
			}

			objChan <- resp
		}
	}()

	return objChan, nil
}
