package geofence

import t38c "github.com/zerobounty/tile38-client"

type BaseObject struct {
	Command string `json:"command"`
	Hook    string `json:"hook,omitempty"`
	Group   string `json:"group"`
	Detect  string `json:"detect"`
	Key     string `json:"key"`
	// TODO: custom time unmarshal
	Time string `json:"time"`
}

// Object struct
type Object struct {
	BaseObject
	ID     string      `json:"id"`
	Object t38c.Object `json:"object"`
}

// Point struct
type Point struct {
	BaseObject
	ID    string     `json:"id"`
	Point t38c.Point `json:"point"`
}

// Bounds struct
type Bounds struct {
	BaseObject
	ID     string      `json:"id"`
	Bounds t38c.Bounds `json:"bounds"`
}

// Hash struct
type Hash struct {
	BaseObject
	ID   string `json:"id"`
	Hash string `json:"hash"`
}

type RoamNearbyFarawayObject struct {
	Key    string      `json:"key"`
	ID     string      `json:"id"`
	Object t38c.Object `json:"object"`
	Meters float64     `json:"meters"`
}

type RoamBaseObject struct {
	BaseObject
	Nearby  *RoamNearbyFarawayObject `json:"nearby,omitempty"`
	Faraway *RoamNearbyFarawayObject `json:"faraway,omitempty"`
}

// RoamObject struct
type RoamObject struct {
	RoamBaseObject
	ID     string             `json:"id"`
	Object t38c.Object        `json:"object"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

// RoamPoint struct
type RoamPoint struct {
	RoamBaseObject
	ID     string             `json:"id"`
	Point  t38c.Point         `json:"point"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

// RoamBounds struct
type RoamBounds struct {
	RoamBaseObject
	ID     string             `json:"id"`
	Bounds t38c.Bounds        `json:"bounds"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

type (
	ObjectChan     chan Object
	PointChan      chan Point
	BoundsChan     chan Bounds
	RoamObjectChan chan RoamObject
	RoamPointChan  chan RoamPoint
	RoamBoundsChan chan Bounds
)

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
