package t38c

import (
	"strconv"

	geojson "github.com/paulmach/orb/geojson"
)

// InwAreaSelector struct
// Intersects Nearby Within
type InwAreaSelector struct {
	client tile38Client
	cmd    string
	key    string
}

func newInwAreaSelector(client tile38Client, cmd, key string) InwAreaSelector {
	return InwAreaSelector{
		client: client,
		cmd:    cmd,
		key:    key,
	}
}

// Get any object that already exists in the database.
func (selector InwAreaSelector) Get(key, objectID string) InwQueryBuilder {
	area := newCmd("GET", key, objectID)
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Bounds - a minimum bounding rectangle.
func (selector InwAreaSelector) Bounds(minlat, minlon, maxlat, maxlon float64) InwQueryBuilder {
	area := newCmd("BOUNDS", floatString(minlat), floatString(minlon), floatString(maxlat), floatString(maxlon))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// FeatureCollection - GeoJSON Feature Collection object.
func (selector InwAreaSelector) FeatureCollection(fc *geojson.FeatureCollection) InwQueryBuilder {
	b, err := fc.MarshalJSON()
	if err != nil {
		panic(err)
	}

	area := newCmd("OBJECT", string(b))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Feature - GeoJSON Feature object.
func (selector InwAreaSelector) Feature(ft *geojson.Feature) InwQueryBuilder {
	b, err := ft.MarshalJSON()
	if err != nil {
		panic(err)
	}

	area := newCmd("OBJECT", string(b))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Geometry - GeoJSON Geometry object.
func (selector InwAreaSelector) Geometry(gm *geojson.Geometry) InwQueryBuilder {
	b, err := gm.MarshalJSON()
	if err != nil {
		panic(err)
	}

	area := newCmd("OBJECT", string(b))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Circle - a circle with the specified center and radius.
func (selector InwAreaSelector) Circle(lat, lon, meters float64) InwQueryBuilder {
	area := newCmd("CIRCLE", floatString(lat), floatString(lon), floatString(meters))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Tile - an XYZ Tile.
func (selector InwAreaSelector) Tile(x, y, z int) InwQueryBuilder {
	area := newCmd("TILE", strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z))
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Quadkey - a QuadKey.
func (selector InwAreaSelector) Quadkey(quadkey string) InwQueryBuilder {
	area := newCmd("QUADKEY", quadkey)
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Hash - a Geohash.
func (selector InwAreaSelector) Hash(hash string) InwQueryBuilder {
	area := newCmd("HASH", hash)
	return newInwQueryBuilder(selector.client, selector.cmd, selector.key, area)
}
