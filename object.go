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

// Object is a Tile38 object.
type Object struct {
	Tile38ID string             `json:"tile38_id"`
	Fields   map[string]float64 `json:"fields"`
}

// GeoJSONObject struct
type GeoJSONObject struct {
	Object
	// Can be geojson.FeatureCollection or geojson.Feature or geojson.Geometry
	GeoJSON interface{} `json:"geojson"`
}

// PointObject struct
type PointObject struct {
	Object
	Point Point `json:"point"`
}

// BoundsObject struct
type BoundsObject struct {
	Object
	Bounds Bounds `json:"bounds"`
}

// HashObject struct
type HashObject struct {
	Object
	Hash string `json:"hash"`
}
