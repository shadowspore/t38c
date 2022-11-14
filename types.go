package t38c

import (
	"strconv"
	"time"

	"github.com/paulmach/orb/geojson"
	"github.com/tidwall/gjson"
)

type field struct {
	Name  string
	Value float64
}

// KeyStats is a tile38 key stats.
type KeyStats struct {
	InMemorySize int `json:"in_memory_size"`
	NumObjects   int `json:"num_objects"`
	NumPoints    int `json:"num_points"`
}

// Point is a tile38 point.
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Bounds is a tile38 bounds object.
type Bounds struct {
	SW Point `json:"sw"`
	NE Point `json:"ne"`
}

// Object is a tile38 object.
type Object struct {
	FeatureCollection *geojson.FeatureCollection `json:"featureCollection,omitempty"`
	Feature           *geojson.Feature           `json:"feature,omitempty"`
	Geometry          *geojson.Geometry          `json:"geometry,omitempty"`
	String            *string                    `json:"string,omitempty"`
}

// UnmarshalJSON unmarshals object from the given json data.
func (ob *Object) UnmarshalJSON(data []byte) error {
	res := gjson.ParseBytes(data)
	objectType := res.Get("type")
	if !objectType.Exists() {
		str := res.String()
		ob.String = &str
		return nil
	}

	var err error
	switch objectType.String() {
	case "FeatureCollection":
		ob.FeatureCollection, err = geojson.UnmarshalFeatureCollection(data)
	case "Feature":
		ob.Feature, err = geojson.UnmarshalFeature(data)
	default:
		ob.Geometry, err = geojson.UnmarshalGeometry(data)
	}

	return err
}

// SearchResponse is a tile38 search response.
type SearchResponse struct {
	Cursor  int      `json:"cursor"`
	Count   int      `json:"count"`
	Fields  []string `json:"fields,omitempty"`
	Objects []struct {
		ID       string    `json:"ID"`
		Object   Object    `json:"object"`
		Fields   []float64 `json:"fields,omitempty"`
		Distance *float64  `json:"distance,omitempty"`
	} `json:"objects,omitempty"`
	Points []struct {
		ID       string    `json:"ID"`
		Point    Point     `json:"point"`
		Fields   []float64 `json:"fields,omitempty"`
		Distance *float64  `json:"distance,omitempty"`
	} `json:"points,omitempty"`
	Bounds []struct {
		ID       string    `json:"ID"`
		Bounds   Bounds    `json:"bounds"`
		Fields   []float64 `json:"fields,omitempty"`
		Distance *float64  `json:"distance,omitempty"`
	} `json:"bounds,omitempty"`
	Hashes []struct {
		ID       string    `json:"id"`
		Hash     string    `json:"hash"`
		Fields   []float64 `json:"fields,omitempty"`
		Distance *float64  `json:"distance,omitempty"`
	} `json:"hashes,omitempty"`
	IDs []string `json:"ids,omitempty"`
}

// OutputFormat specifies expected format.
type OutputFormat cmd

var (
	// FormatCount - Total object count sent in the response.
	// When LIMIT or CURSOR are provided, COUNT returns the number of results that would otherwise be sent as objects.
	// When LIMIT is not specified, COUNT totals up all items starting from provided CURSOR position
	// (or zero if a cursor is omitted). LIMIT and CURSOR options are ignored.
	FormatCount = OutputFormat(newCmd("COUNT"))
	// FormatIDs - A list of IDs belonging to the key. Will not return the objects.
	FormatIDs = OutputFormat(newCmd("IDS"))
	// FormatPoints - A list of standard latitude, longitude points.
	FormatPoints = OutputFormat(newCmd("POINTS"))
	// FormatBounds - A list of minimum bounding rectangle.
	FormatBounds = OutputFormat(newCmd("BOUNDS"))
	// FormatHashes - A list of Geohash. Requires a precision of 1 to 22.
	FormatHashes = func(precision int) OutputFormat {
		return OutputFormat(newCmd("HASHES", strconv.Itoa(precision)))
	}
)

// Meta is tile38 metadata.
type Meta struct {
	Name  string
	Value string
}

// Hook represents tile38 channel.
type Hook struct {
	Endpoints []string `json:"endpoints"`
	Chan
}

// Chan represents tile38 channel.
type Chan struct {
	Name    string            `json:"name"`
	Key     string            `json:"key"`
	Command []string          `json:"command"`
	Meta    map[string]string `json:"meta"`
}

// EventHandler handles tile38 events.
type EventHandler interface {
	// HandleEvent handles tile38 event.
	HandleEvent(event *GeofenceEvent) error
}

// EventHandlerFunc is an adapter to allow the use of
// ordinary functions as tile38 event handlers.
type EventHandlerFunc func(event *GeofenceEvent) error

// HandleEvent handles tile38 event.
func (e EventHandlerFunc) HandleEvent(event *GeofenceEvent) error {
	return e(event)
}

// GeofenceEvent is a tile38 geofence event.
type GeofenceEvent struct {
	Command string             `json:"command"`
	Hook    string             `json:"hook,omitempty"`
	Group   string             `json:"group"`
	Detect  string             `json:"detect"`
	Key     string             `json:"key"`
	Time    time.Time          `json:"time"`
	ID      string             `json:"id"`
	Object  *Object            `json:"object,omitempty"`
	Point   *Point             `json:"point,omitempty"`
	Bounds  *Bounds            `json:"bounds,omitempty"`
	Hash    *string            `json:"hash,omitempty"`
	Nearby  *RoamObject        `json:"nearby,omitempty"`
	Faraway *RoamObject        `json:"faraway,omitempty"`
	Fields  map[string]float64 `json:"fields,omitempty"`
}

// RoamObject is a tile38 roam object.
type RoamObject struct {
	Key    string  `json:"key"`
	ID     string  `json:"id"`
	Object Object  `json:"object"`
	Meters float64 `json:"meters"`
}

// NotifyCommand ...
type NotifyCommand string

const (
	// Del notifies the client that an object has been deleted from the collection that is being fenced.
	Del NotifyCommand = "del"
	// Drop notifies the client that the entire collection is dropped.
	Drop NotifyCommand = "drop"
	// Set notifies the client that an object has been added or updated,
	// and when it’s position is detected by the fence.
	Set NotifyCommand = "set"
)

// DetectAction ...
type DetectAction string

const (
	// Inside is when an object is inside the specified area.
	Inside DetectAction = "inside"
	// Outside is when an object is outside the specified area.
	Outside DetectAction = "outside"
	// Enter is when an object that was not previously in the fence has entered the area.
	Enter DetectAction = "enter"
	// Exit is when an object that was previously in the fence has exited the area.
	Exit DetectAction = "exit"
	// Cross is when an object that was not previously in the fence has entered and exited the area.
	Cross DetectAction = "cross"
)
