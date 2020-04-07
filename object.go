package t38c

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

// BaseObject is a Tile38 object.
type BaseObject struct {
	Tile38ID string             `json:"tile38_id"`
	Fields   map[string]float64 `json:"fields"`
	Distance *float64           `json:"distance,omitempty"`
}

// Object struct
type Object struct {
	BaseObject
	// Can be geojson.FeatureCollection or geojson.Feature or geojson.Geometry or string
	Object interface{} `json:"object"`
}

// PointObject struct
type PointObject struct {
	BaseObject
	Point Point `json:"point"`
}

// BoundsObject struct
type BoundsObject struct {
	BaseObject
	Bounds Bounds `json:"bounds"`
}

// HashObject struct
type HashObject struct {
	BaseObject
	Hash string `json:"hash"`
}
