package t38c

import (
	"strconv"

	geojson "github.com/paulmach/orb/geojson"
)

// GeofenceAreaSelector struct
type GeofenceAreaSelector struct {
	client tile38Client
	cmd    string
	key    string
}

func newGeofenceAreaSelector(client tile38Client, cmd, key string) GeofenceAreaSelector {
	return GeofenceAreaSelector{
		client: client,
		cmd:    cmd,
		key:    key,
	}
}

// Get any object that already exists in the database.
func (selector GeofenceAreaSelector) Get(key, objectID string) GeofenceQueryBuilder {
	area := newCmd("GET", key, objectID)
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Bounds - a minimum bounding rectangle.
func (selector GeofenceAreaSelector) Bounds(minlat, minlon, maxlat, maxlon float64) GeofenceQueryBuilder {
	area := newCmd("BOUNDS", floatString(minlat), floatString(minlon), floatString(maxlat), floatString(maxlon))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// FeatureCollection - GeoJSON Feature Collection object.
func (selector GeofenceAreaSelector) FeatureCollection(fc *geojson.FeatureCollection) GeofenceQueryBuilder {
	b, err := fc.MarshalJSON()
	if err != nil {
		panic(err)
	}

	area := newCmd("OBJECT", string(b))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Feature - GeoJSON Feature object.
func (selector GeofenceAreaSelector) Feature(ft *geojson.Feature) GeofenceQueryBuilder {
	b, err := ft.MarshalJSON()
	if err != nil {
		panic(err)
	}

	area := newCmd("OBJECT", string(b))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Geometry - GeoJSON Geometry object.
func (selector GeofenceAreaSelector) Geometry(gm *geojson.Geometry) GeofenceQueryBuilder {
	b, err := gm.MarshalJSON()
	if err != nil {
		panic(err)
	}

	area := newCmd("OBJECT", string(b))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Circle - a circle with the specified center and radius.
func (selector GeofenceAreaSelector) Circle(lat, lon, meters float64) GeofenceQueryBuilder {
	area := newCmd("CIRCLE", floatString(lat), floatString(lon), floatString(meters))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Tile - an XYZ Tile.
func (selector GeofenceAreaSelector) Tile(x, y, z int) GeofenceQueryBuilder {
	area := newCmd("TILE", strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z))
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Quadkey - a QuadKey.
func (selector GeofenceAreaSelector) Quadkey(quadkey string) GeofenceQueryBuilder {
	area := newCmd("QUADKEY", quadkey)
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Hash - a Geohash.
func (selector GeofenceAreaSelector) Hash(hash string) GeofenceQueryBuilder {
	area := newCmd("HASH", hash)
	return newGeofenceQueryBuilder(selector.client, selector.cmd, selector.key, area)
}
