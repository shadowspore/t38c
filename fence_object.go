package t38c

// GeofenceBaseObject struct
type GeofenceBaseObject struct {
	Command string `json:"command"`
	Hook    string `json:"hook"`
	Group   string `json:"group"`
	Detect  string `json:"detect"`
	Key     string `json:"key"`
	Time    string `json:"time"`
}

type GeofenceRoamNearbyFarawayObject struct {
	Key    string  `json:"key"`
	ID     string  `json:"id"`
	Object Object  `json:"object"`
	Meters float64 `json:"meters"`
}

type GeofenceRoamBaseObject struct {
	GeofenceBaseObject
	Nearby  *GeofenceRoamNearbyFarawayObject `json:"nearby,omitempty"`
	Faraway *GeofenceRoamNearbyFarawayObject `json:"faraway,omitempty"`
}

// GeofenceRoamObject struct
type GeofenceRoamObject struct {
	GeofenceRoamBaseObject
	Object Object `json:"object"`
}

// GeofenceRoamPoint struct
type GeofenceRoamPoint struct {
	GeofenceRoamBaseObject
	Point Point `json:"point"`
}
