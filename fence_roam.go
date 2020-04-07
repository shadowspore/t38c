package t38c

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func buildFenceArgs(objectType, key, target, pattern string, meters int, detectActions []FenceDetectAction, opts []SearchOption) []string {
	args := []string{
		key, "FENCE",
	}

	if len(detectActions) > 0 {
		args = append(args, "DETECT")
		actions := ""
		first := true
		for _, action := range detectActions {
			if !first {
				actions += ","
			}
			actions += string(action)
			first = false
		}
		args = append(args, actions)
	}

	if len(objectType) > 0 {
		args = append(args, objectType)
	}

	args = append(args, []string{
		"ROAM", target, pattern, strconv.Itoa(meters),
	}...)

	for _, opt := range opts {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}

	return args
}

type GeofenceRoamObjectChan chan GeofenceRoamObject

// FenceRoam ...
func (client *FenceClient) FenceRoam(key, target, pattern string, meters int, detectActions []FenceDetectAction, opts ...SearchOption) (GeofenceRoamObjectChan, error) {
	args := buildFenceArgs("", key, target, pattern, meters, detectActions, opts)

	ch, err := client.exec.Fence("NEARBY", args...)
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

			var resp GeofenceRoamObject
			if err := json.Unmarshal(event.Data, &resp); err != nil {
				panic(err)
			}

			objChan <- resp
		}
	}()

	return objChan, nil
}

type GeofenceRoamPointChan chan GeofenceRoamPoint

func (client *FenceClient) FenceRoamPoints(key, target, pattern string, meters int, detectActions []FenceDetectAction, opts ...SearchOption) (GeofenceRoamPointChan, error) {
	args := buildFenceArgs("POINTS", key, target, pattern, meters, detectActions, opts)

	ch, err := client.exec.Fence("NEARBY", args...)
	if err != nil {
		return nil, err
	}

	objChan := make(GeofenceRoamPointChan, 10)
	go func() {
		defer close(objChan)
		for event := range ch {
			if event.Err != nil {
				log.Fatal(event.Err)
			}
			fmt.Printf("%s\n\n\n", event.Data)
			var resp GeofenceRoamPoint
			if err := json.Unmarshal(event.Data, &resp); err != nil {
				panic(err)
			}

			objChan <- resp
		}
	}()

	return objChan, nil
}

type GeofenceRoamIDsChan chan GeofenceRoamID

func (client *FenceClient) FenceRoamIDs(key, target, pattern string, meters int, detectActions []FenceDetectAction, opts ...SearchOption) (GeofenceRoamIDsChan, error) {
	args := buildFenceArgs("IDS", key, target, pattern, meters, detectActions, opts)

	ch, err := client.exec.Fence("NEARBY", args...)
	if err != nil {
		return nil, err
	}

	objChan := make(GeofenceRoamIDsChan, 10)
	go func() {
		defer close(objChan)
		for event := range ch {
			if event.Err != nil {
				log.Fatal(event.Err)
			}
			fmt.Printf("%s\n\n\n", event.Data)
			var resp GeofenceRoamID
			if err := json.Unmarshal(event.Data, &resp); err != nil {
				panic(err)
			}

			objChan <- resp
		}
	}()

	return objChan, nil
}
