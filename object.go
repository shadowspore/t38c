package t38c

import (
	geojson "github.com/paulmach/go.geojson"
	"github.com/tidwall/gjson"
)

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

// Object is a Tile38 object.
type Object struct {
	Tile38ID string `json:"tile38_id"` // get this
	// Can be geojson.FeatureCollection or geojson.Feature or geojson.Geometry
	GeoJSON interface{}    `json:"geojson"`
	Fields  map[string]int `json:"fields"` // and this from up-level json
}

// UnmarshalJSON decodes the data into a GeoJSON object.
func (ob *Object) UnmarshalJSON(data []byte) (err error) {
	ob.GeoJSON, err = unmarshalGeoJSON(data)
	return
}

func unmarshalGeoJSON(data []byte) (interface{}, error) {
	switch gjson.GetBytes(data, "type").String() {
	case "FeatureCollection":
		return geojson.UnmarshalFeatureCollection(data)
	case "Feature":
		return geojson.UnmarshalFeature(data)
	default:
		return geojson.UnmarshalGeometry(data)
	}
}
