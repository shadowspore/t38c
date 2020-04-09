package t38c

type geofenceBaseObject struct {
	Command string `json:"command"`
	Hook    string `json:"hook,omitempty"`
	Group   string `json:"group"`
	Detect  string `json:"detect"`
	Key     string `json:"key"`
	// TODO: custom time unmarshal
	Time string `json:"time"`
}

// GeofenceObject struct
type GeofenceObject struct {
	geofenceBaseObject
	ID     string `json:"id"`
	Object Object `json:"object"`
}

// GeofencePoint struct
type GeofencePoint struct {
	geofenceBaseObject
	ID    string `json:"id"`
	Point Point  `json:"point"`
}

// GeofenceBounds struct
type GeofenceBounds struct {
	geofenceBaseObject
	Bounds Bounds `json:"bounds"`
	Object Object `json:"object"`
}

// GeofenceHash struct
type GeofenceHash struct {
	geofenceBaseObject
	Hash   string `json:"hash"`
	Object Object `json:"object"`
}

type geofenceRoamNearbyFarawayObject struct {
	Key    string  `json:"key"`
	ID     string  `json:"id"`
	Object Object  `json:"object"`
	Meters float64 `json:"meters"`
}

type geofenceRoamBaseObject struct {
	geofenceBaseObject
	Nearby  *geofenceRoamNearbyFarawayObject `json:"nearby,omitempty"`
	Faraway *geofenceRoamNearbyFarawayObject `json:"faraway,omitempty"`
}

// GeofenceRoamObject struct
type GeofenceRoamObject struct {
	geofenceRoamBaseObject
	ID     string             `json:"id"`
	Object Object             `json:"object"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

// GeofenceRoamPoint struct
type GeofenceRoamPoint struct {
	geofenceRoamBaseObject
	ID     string             `json:"id"`
	Point  Point              `json:"point"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

// GeofenceRoamBounds struct
type GeofenceRoamBounds struct {
	geofenceRoamBaseObject
	ID     string             `json:"id"`
	Bounds Bounds             `json:"bounds"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

type (
	GeofenceObjectChan     chan GeofenceObject
	GeofencePointChan      chan GeofencePoint
	GeofenceBoundsChan     chan GeofenceBounds
	GeofenceRoamObjectChan chan GeofenceRoamObject
	GeofenceRoamPointChan  chan GeofenceRoamPoint
	GeofenceRoamBoundsChan chan GeofenceBounds
)

// GeofenceDetectAction ...
type GeofenceDetectAction string

const (
	// Inside action
	Inside GeofenceDetectAction = "inside"
	// Outside action
	Outside GeofenceDetectAction = "outside"
	// Enter action
	Enter GeofenceDetectAction = "enter"
	// Exit action
	Exit GeofenceDetectAction = "exit"
	// Cross action
	Cross GeofenceDetectAction = "cross"
)
