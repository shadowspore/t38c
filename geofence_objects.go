package t38c

type geofenceBaseObject struct {
	Command string `json:"command"`
	Hook    string `json:"hook"`
	Group   string `json:"group"`
	Detect  string `json:"detect"`
	Key     string `json:"key"`
	Time    string `json:"time"`
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

// GeofenceRoamObjectChan ...
type GeofenceRoamObjectChan chan GeofenceRoamObject

// GeofenceRoamObject struct
type GeofenceRoamObject struct {
	geofenceRoamBaseObject
	Object Object             `json:"object"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

// GeofenceRoamPointChan ...
type GeofenceRoamPointChan chan GeofenceRoamPoint

// GeofenceRoamPoint struct
type GeofenceRoamPoint struct {
	geofenceRoamBaseObject
	Point  Point              `json:"point"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

// GeofenceRoamBoundsChan ...
type GeofenceRoamBoundsChan chan GeofenceRoamBounds

// GeofenceRoamBounds struct
type GeofenceRoamBounds struct {
	geofenceRoamBaseObject
	Bounds Bounds             `json:"bounds"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

type GeofenceDetectAction string

const (
	Inside  GeofenceDetectAction = "inside"
	Outside GeofenceDetectAction = "outside"
	Enter   GeofenceDetectAction = "enter"
	Exit    GeofenceDetectAction = "exit"
	Cross   GeofenceDetectAction = "cross"
)

func Detect(actions ...GeofenceDetectAction) []GeofenceDetectAction {
	return actions
}
