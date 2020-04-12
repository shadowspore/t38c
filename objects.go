package t38c

import (
	"strconv"

	geojson "github.com/paulmach/go.geojson"
	"github.com/tidwall/gjson"
)

// KeyStats struct
type KeyStats struct {
	InMemorySize int `json:"in_memory_size"`
	NumObjects   int `json:"num_objects"`
	NumPoints    int `json:"num_points"`
}

// Point struct
type Point struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

// Bounds struct
type Bounds struct {
	SW Point `json:"sw"`
	NE Point `json:"ne"`
}

// Object struct
type Object struct {
	FeatureCollection *geojson.FeatureCollection `json:"featureCollection,omitempty"`
	Feature           *geojson.Feature           `json:"feature,omitempty"`
	Geometry          *geojson.Geometry          `json:"geometry,omitempty"`
	String            *string                    `json:"string,omitempty"`
}

// UnmarshalJSON ...
func (ob *Object) UnmarshalJSON(data []byte) error {
	res := gjson.ParseBytes(data)
	objectType := res.Get("type")
	if !objectType.Exists() {
		str := res.String()
		ob.String = &str
	} else {
		var err error
		switch objectType.String() {
		case "FeatureCollection":
			ob.FeatureCollection, err = geojson.UnmarshalFeatureCollection(data)
		case "Feature":
			ob.Feature, err = geojson.UnmarshalFeature(data)
		default:
			ob.Geometry, err = geojson.UnmarshalGeometry(data)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// GetResponse struct
type GetResponse struct {
	Object *Object            `json:"object,omitempty"`
	Point  *Point             `json:"point,omitempty"`
	Bounds *Bounds            `json:"bounds,omitempty"`
	Hash   *string            `json:"hash,omitempty"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

// SearchResponse struct
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

// OutputFormat ...
type OutputFormat Command

var (
	// FormatCount ...
	FormatCount = OutputFormat(NewCommand("COUNT"))
	// FormatIDs ...
	FormatIDs = OutputFormat(NewCommand("IDS"))
	// FormatPoints ...
	FormatPoints = OutputFormat(NewCommand("POINTS"))
	// FormatBounds ...
	FormatBounds = OutputFormat(NewCommand("BOUNDS"))
	// FormatHashes ...
	FormatHashes = func(precision int) OutputFormat {
		return OutputFormat(NewCommand("HASHES", strconv.Itoa(precision)))
	}
)
