package t38c

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

// GeofenceBaseObject struct
type GeofenceBaseObject struct {
	Command string `json:"command"`
	Hook    string `json:"hook"`
	Group   string `json:"group"`
	Detect  string `json:"detect"`
	Key     string `json:"key"`
	Time    string `json:"time"`
	ID      string `json:"id"`
}

type GeofenceRoamNearbyFarawayObject struct {
	Key    string      `json:"key"`
	ID     string      `json:"id"`
	Object interface{} `json:"object"`
	Meters float64     `json:"meters"`
}

type GeofenceRoamBaseObject struct {
	GeofenceBaseObject
	Nearby  *GeofenceRoamNearbyFarawayObject `json:"nearby,omitempty"`
	Faraway *GeofenceRoamNearbyFarawayObject `json:"faraway,omitempty"`
}

func (gf *GeofenceRoamBaseObject) UnmarshalJSON(data []byte) error {
	var resp struct {
		GeofenceBaseObject
		Nearby *struct {
			Key    string          `json:"key"`
			ID     string          `json:"id"`
			Object json.RawMessage `json:"object"`
			Meters float64         `json:"meters"`
		} `json:"nearby,omitempty"`
		Faraway *struct {
			Key    string          `json:"key"`
			ID     string          `json:"id"`
			Object json.RawMessage `json:"object"`
			Meters float64         `json:"meters"`
		} `json:"faraway,omitempty"`
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}

	gf.GeofenceBaseObject = resp.GeofenceBaseObject

	if resp.Nearby != nil {
		nearbyObj, err := unmarshalObject(resp.Nearby.Object)
		if err != nil {
			return err
		}

		gf.Nearby = &GeofenceRoamNearbyFarawayObject{
			Key:    resp.Nearby.Key,
			ID:     resp.Nearby.ID,
			Object: nearbyObj,
			Meters: resp.Nearby.Meters,
		}
	}

	if resp.Faraway != nil {
		farawayObj, err := unmarshalObject(resp.Faraway.Object)
		if err != nil {
			return err
		}

		gf.Faraway = &GeofenceRoamNearbyFarawayObject{
			Key:    resp.Faraway.Key,
			ID:     resp.Faraway.ID,
			Object: farawayObj,
			Meters: resp.Faraway.Meters,
		}
	}

	return nil
}

// GeofenceRoamObject struct
type GeofenceRoamObject struct {
	GeofenceRoamBaseObject
	Object interface{} `json:"object"`
}

// UnmarshalJSON ...
func (gf *GeofenceRoamObject) UnmarshalJSON(data []byte) (err error) {
	var base GeofenceRoamBaseObject
	if err := json.Unmarshal(data, &base); err != nil {
		return err
	}

	gf.GeofenceRoamBaseObject = base
	obj, err := unmarshalObject([]byte(gjson.GetBytes(data, "object").String()))
	if err != nil {
		return err
	}

	gf.Object = obj
	return nil
}

// GeofenceRoamPoint struct
type GeofenceRoamPoint struct {
	GeofenceRoamBaseObject
	Point Point `json:"point"`
} 
