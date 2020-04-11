package geofence

import t38c "github.com/zerobounty/tile38-client"

type Response struct {
	Command string `json:"command"`
	Hook    string `json:"hook,omitempty"`
	Group   string `json:"group"`
	Detect  string `json:"detect"`
	Key     string `json:"key"`
	// TODO: custom time unmarshal
	Time    string                   `json:"time"`
	ID      string                   `json:"id"`
	Object  *t38c.Object             `json:"object,omitempty"`
	Point   *t38c.Point              `json:"point,omitempty"`
	Bounds  *t38c.Bounds             `json:"bounds,omitempty"`
	Hash    *string                  `json:"hash,omitempty"`
	Nearby  *roamNearbyFarawayObject `json:"nearby,omitempty"`
	Faraway *roamNearbyFarawayObject `json:"faraway,omitempty"`
	Fields  map[string]float64       `json:"fields,omitempty"`
}

type roamNearbyFarawayObject struct {
	Key    string      `json:"key"`
	ID     string      `json:"id"`
	Object t38c.Object `json:"object"`
	Meters float64     `json:"meters"`
}

// DetectAction ...
type DetectAction string

const (
	// Inside action
	Inside DetectAction = "inside"
	// Outside action
	Outside DetectAction = "outside"
	// Enter action
	Enter DetectAction = "enter"
	// Exit action
	Exit DetectAction = "exit"
	// Cross action
	Cross DetectAction = "cross"
)
