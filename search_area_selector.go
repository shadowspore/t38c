package t38c

import (
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

// SearchAreaSelector struct
type SearchAreaSelector struct {
	client *Client
	cmd    string
	key    string
}

func newSearchAreaSelector(client *Client, cmd, key string) SearchAreaSelector {
	return SearchAreaSelector{
		client: client,
		cmd:    cmd,
		key:    key,
	}
}

// Get any object that already exists in the database.
func (selector SearchAreaSelector) Get(objectID string) SearchQueryBuilder {
	area := NewCommand("GET", objectID)
	return newSearchQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Bounds - a minimum bounding rectangle.
func (selector SearchAreaSelector) Bounds(minlat, minlon, maxlat, maxlon float64) SearchQueryBuilder {
	area := NewCommand("BOUNDS", floatString(minlat), floatString(minlon), floatString(maxlat), floatString(maxlon))
	return newSearchQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// FeatureCollection - GeoJSON Feature Collection object.
func (selector SearchAreaSelector) FeatureCollection(fc *geojson.FeatureCollection) SearchQueryBuilder {
	// TODO: handle error?
	b, _ := fc.MarshalJSON()
	area := NewCommand("OBJECT", string(b))
	return newSearchQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Feature - GeoJSON Feature object.
func (selector SearchAreaSelector) Feature(ft *geojson.Feature) SearchQueryBuilder {
	// TODO: handle error?
	b, _ := ft.MarshalJSON()
	area := NewCommand("OBJECT", string(b))
	return newSearchQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Geometry - GeoJSON Geometry object.
func (selector SearchAreaSelector) Geometry(gm *geojson.Geometry) SearchQueryBuilder {
	// TODO: handle error?
	b, _ := gm.MarshalJSON()
	area := NewCommand("OBJECT", string(b))
	return newSearchQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Circle - a circle with the specified center and radius.
func (selector SearchAreaSelector) Circle(lat, lon, meters float64) SearchQueryBuilder {
	area := NewCommand("CIRCLE", floatString(lat), floatString(lon), floatString(meters))
	return newSearchQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Tile - an XYZ Tile.
func (selector SearchAreaSelector) Tile(x, y, z int) SearchQueryBuilder {
	area := NewCommand("TILE", strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z))
	return newSearchQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Quadkey - a QuadKey.
func (selector SearchAreaSelector) Quadkey(quadkey string) SearchQueryBuilder {
	area := NewCommand("QUADKEY", quadkey)
	return newSearchQueryBuilder(selector.client, selector.cmd, selector.key, area)
}

// Hash - a Geohash.
func (selector SearchAreaSelector) Hash(hash string) SearchQueryBuilder {
	area := NewCommand("HASH", hash)
	return newSearchQueryBuilder(selector.client, selector.cmd, selector.key, area)
}
