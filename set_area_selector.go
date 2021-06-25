package t38c

import geojson "github.com/paulmach/orb/geojson"

// SetAreaSelector struct
type SetAreaSelector struct {
	client   tile38Client
	key      string
	objectID string
}

func newSetAreaSelector(client tile38Client, key, objectID string) SetAreaSelector {
	return SetAreaSelector{
		client:   client,
		key:      key,
		objectID: objectID,
	}
}

// Point set a simple point in latitude, longitude.
func (selector SetAreaSelector) Point(lat, lon float64) SetQueryBuilder {
	area := newCmd("POINT", floatString(lat), floatString(lon))
	return newSetQueryBuilder(selector.client, selector.key, selector.objectID, area)
}

// PointZ - a point with Z coordinate.
// This is application specific such as elevation, or a timestamp, etc.
func (selector SetAreaSelector) PointZ(lat, lon, z float64) SetQueryBuilder {
	area := newCmd("POINT", floatString(lat), floatString(lon), floatString(z))
	return newSetQueryBuilder(selector.client, selector.key, selector.objectID, area)
}

// Bounds - a bounding box consists of two points.
// The first being the southwestern most point and the second is the northeastern most point.
func (selector SetAreaSelector) Bounds(lat1, lon1, lat2, lon2 float64) SetQueryBuilder {
	area := newCmd("BOUNDS", floatString(lat1), floatString(lon1), floatString(lat2), floatString(lon2))
	return newSetQueryBuilder(selector.client, selector.key, selector.objectID, area)
}

// FeatureCollection - set GeoJSON Feature Collection object.
func (selector SetAreaSelector) FeatureCollection(fc *geojson.FeatureCollection) SetQueryBuilder {
	b, _ := fc.MarshalJSON()
	area := newCmd("OBJECT", string(b))
	return newSetQueryBuilder(selector.client, selector.key, selector.objectID, area)
}

// Feature - set GeoJSON Feature object.
func (selector SetAreaSelector) Feature(ft *geojson.Feature) SetQueryBuilder {
	b, _ := ft.MarshalJSON()
	area := newCmd("OBJECT", string(b))
	return newSetQueryBuilder(selector.client, selector.key, selector.objectID, area)
}

// Geometry - set GeoJSON Geometry object.
func (selector SetAreaSelector) Geometry(gm *geojson.Geometry) SetQueryBuilder {
	b, _ := gm.MarshalJSON()
	area := newCmd("OBJECT", string(b))
	return newSetQueryBuilder(selector.client, selector.key, selector.objectID, area)
}

// Hash - A geohash is a convenient way of expressing a location (anywhere in the world)
// using a short alphanumeric string, with greater precision obtained with longer strings.
func (selector SetAreaSelector) Hash(hash string) SetQueryBuilder {
	area := newCmd("HASH", hash)
	return newSetQueryBuilder(selector.client, selector.key, selector.objectID, area)
}

// String - It’s possible to set a raw string.
// The value of a string type can be plain text or a series of raw bytes.
// To retrieve a string value you can use GET, SCAN, or SEARCH.
func (selector SetAreaSelector) String(str string) SetQueryBuilder {
	area := newCmd("STRING", str)
	return newSetQueryBuilder(selector.client, selector.key, selector.objectID, area)
}
