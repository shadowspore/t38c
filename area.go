package t38c

import (
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

// SearchArea ...
type SearchArea Command

var (
	// Get any object that already exists in the database.
	Get = func(objectID string) SearchArea {
		return SearchArea(NewCommand("GET", objectID))
	}

	// AreaBounds - a minimum bounding rectangle.
	AreaBounds = func(minlat, minlon, maxlat, maxlon float64) SearchArea {
		return SearchArea(NewCommand("BOUNDS", floatString(minlat), floatString(minlon), floatString(maxlat), floatString(maxlon)))
	}

	// AreaFeatureCollection - GeoJSON Feature Collection object.
	AreaFeatureCollection = func(fc *geojson.FeatureCollection) SearchArea {
		// TODO: handle error?
		b, _ := fc.MarshalJSON()
		return SearchArea(NewCommand("OBJECT", string(b)))
	}

	// AreaFeature - GeoJSON Feature object.
	AreaFeature = func(ft *geojson.Feature) SearchArea {
		// TODO: handle error?
		b, _ := ft.MarshalJSON()
		return SearchArea(NewCommand("OBJECT", string(b)))
	}

	// AreaGeometry - GeoJSON Geometry object.
	AreaGeometry = func(gm *geojson.Geometry) SearchArea {
		// TODO: handle error?
		b, _ := gm.MarshalJSON()
		return SearchArea(NewCommand("OBJECT", string(b)))
	}

	// AreaCircle - a circle with the specified center and radius.
	AreaCircle = func(lat, lon, meters float64) SearchArea {
		return SearchArea(NewCommand("CIRCLE", floatString(lat), floatString(lon), floatString(meters)))
	}

	// AreaTile - an XYZ Tile.
	AreaTile = func(x, y, z int) SearchArea {
		return SearchArea(NewCommand("TILE", strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z)))
	}

	// AreaQuadkey - a QuadKey.
	AreaQuadkey = func(quadkey string) SearchArea {
		return SearchArea(NewCommand("QUADKEY", quadkey))
	}

	// AreaHash - a Geohash.
	AreaHash = func(hash string) SearchArea {
		return SearchArea(NewCommand("HASH", hash))
	}
)

// SetArea ...
type SetArea Command

var (
	// SetPoint set a simple point in latitude, longitude.
	SetPoint = func(lat, lon float64) SetArea {
		return SetArea(NewCommand("POINT", floatString(lat), floatString(lon)))
	}

	// SetPointZ - a point with Z coordinate.
	// This is application specific such as elevation, or a timestamp, etc.
	SetPointZ = func(lat, lon, z float64) SetArea {
		return SetArea(NewCommand("POINT", floatString(lat), floatString(lon), floatString(z)))
	}

	// SetBounds - a bounding box consists of two points.
	// The first being the southwestern most point and the second is the northeastern most point.
	SetBounds = func(lat1, lon1, lat2, lon2 float64) SetArea {
		return SetArea(NewCommand("BOUNDS", floatString(lat1), floatString(lon1), floatString(lat2), floatString(lon2)))
	}

	// SetFeatureCollection - set GeoJSON Feature Collection object.
	SetFeatureCollection = func(fc *geojson.FeatureCollection) SetArea {
		b, _ := fc.MarshalJSON()
		return SetArea(NewCommand("OBJECT", string(b)))
	}

	// SetFeature - set GeoJSON Feature object.
	SetFeature = func(ft *geojson.Feature) SetArea {
		b, _ := ft.MarshalJSON()
		return SetArea(NewCommand("OBJECT", string(b)))
	}

	// SetGeometry - set GeoJSON Geometry object.
	SetGeometry = func(gm *geojson.Geometry) SetArea {
		b, _ := gm.MarshalJSON()
		return SetArea(NewCommand("OBJECT", string(b)))
	}

	// SetHash - A geohash is a convenient way of expressing a location (anywhere in the world)
	// using a short alphanumeric string, with greater precision obtained with longer strings.
	SetHash = func(hash string) SetArea {
		return SetArea(NewCommand("HASH", hash))
	}

	// SetString - It’s possible to set a raw string.
	// The value of a string type can be plain text or a series of raw bytes.
	// To retrieve a string value you can use GET, SCAN, or SEARCH.
	SetString = func(str string) SetArea {
		return SetArea(NewCommand("STRING", str))
	}
)
