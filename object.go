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

// GetObjectResponse struct
type GetObjectResponse struct {
	Object Object             `json:"object"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

// GetPointResponse struct
type GetPointResponse struct {
	Point  Point              `json:"point"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

// GetBoundsResponse struct
type GetBoundsResponse struct {
	Bounds Bounds             `json:"bounds"`
	Fields map[string]float64 `json:"fields,omitempty"`
}

// GetHashResponse struct
type GetHashResponse struct {
	Hash   string             `json:"hash"`
	Fields map[string]float64 `json:"fields,omitempty"`
}
