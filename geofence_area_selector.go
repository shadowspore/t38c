package t38c

import (
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

// GeofenceAreaSelector struct
type GeofenceAreaSelector struct {
	client *Client
	cmd    string
	key    string
}

func newGeofenceAreaSelector(client *Client, cmd, key string) GeofenceAreaSelector {
	return GeofenceAreaSelector{
		client: client,
		cmd:    cmd,
		key:    key,
	}
}

// Get any object that already exists in the database.
func (selector GeofenceAreaSelector) Get(objectID string) GeofenceQueryBuilder {
	area := NewCommand("GET", objectID)
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Bounds - a minimum bounding rectangle.
func (selector GeofenceAreaSelector) Bounds(minlat, minlon, maxlat, maxlon float64) GeofenceQueryBuilder {
	area := NewCommand("BOUNDS", floatString(minlat), floatString(minlon), floatString(maxlat), floatString(maxlon))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// FeatureCollection - GeoJSON Feature Collection object.
func (selector GeofenceAreaSelector) FeatureCollection(fc *geojson.FeatureCollection) GeofenceQueryBuilder {
	// TODO: handle error?
	b, _ := fc.MarshalJSON()
	area := NewCommand("OBJECT", string(b))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Feature - GeoJSON Feature object.
func (selector GeofenceAreaSelector) Feature(ft *geojson.Feature) GeofenceQueryBuilder {
	// TODO: handle error?
	b, _ := ft.MarshalJSON()
	area := NewCommand("OBJECT", string(b))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Geometry - GeoJSON Geometry object.
func (selector GeofenceAreaSelector) Geometry(gm *geojson.Geometry) GeofenceQueryBuilder {
	// TODO: handle error?
	b, _ := gm.MarshalJSON()
	area := NewCommand("OBJECT", string(b))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Circle - a circle with the specified center and radius.
func (selector GeofenceAreaSelector) Circle(lat, lon, meters float64) GeofenceQueryBuilder {
	area := NewCommand("CIRCLE", floatString(lat), floatString(lon), floatString(meters))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Tile - an XYZ Tile.
func (selector GeofenceAreaSelector) Tile(x, y, z int) GeofenceQueryBuilder {
	area := NewCommand("TILE", strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Quadkey - a QuadKey.
func (selector GeofenceAreaSelector) Quadkey(quadkey string) GeofenceQueryBuilder {
	area := NewCommand("QUADKEY", quadkey)
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Hash - a Geohash.
func (selector GeofenceAreaSelector) Hash(hash string) GeofenceQueryBuilder {
	area := NewCommand("HASH", hash)
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}
