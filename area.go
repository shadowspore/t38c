package t38c

import (
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

// SearchArea ...
type SearchArea Command

var (
	// Get ...
	Get = func(objectID string) SearchArea {
		return SearchArea(NewCommand("GET", objectID))
	}

	// AreaBounds ...
	AreaBounds = func(minlat, minlon, maxlat, maxlon float64) SearchArea {
		return SearchArea(NewCommand("BOUNDS", floatString(minlat), floatString(minlon), floatString(maxlat), floatString(maxlon)))
	}

	// AreaFeatureCollection ...
	AreaFeatureCollection = func(fc *geojson.FeatureCollection) SearchArea {
		// TODO: handle error?
		b, _ := fc.MarshalJSON()
		return SearchArea(NewCommand("OBJECT", string(b)))
	}

	// AreaFeature ...
	AreaFeature = func(ft *geojson.Feature) SearchArea {
		// TODO: handle error?
		b, _ := ft.MarshalJSON()
		return SearchArea(NewCommand("OBJECT", string(b)))
	}

	// AreaGeometry ...
	AreaGeometry = func(gm *geojson.Geometry) SearchArea {
		// TODO: handle error?
		b, _ := gm.MarshalJSON()
		return SearchArea(NewCommand("OBJECT", string(b)))
	}

	// AreaCircle ...
	AreaCircle = func(lat, lon, meters float64) SearchArea {
		return SearchArea(NewCommand("CIRCLE", floatString(lat), floatString(lon), floatString(meters)))
	}

	// AreaTile ...
	AreaTile = func(x, y, z int) SearchArea {
		return SearchArea(NewCommand("TILE", strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z)))
	}

	// AreaQuadkey ...
	AreaQuadkey = func(quadkey string) SearchArea {
		return SearchArea(NewCommand("QUADKEY", quadkey))
	}

	// AreaHash ...
	AreaHash = func(hash string) SearchArea {
		return SearchArea(NewCommand("HASH", hash))
	}
)

// SetArea ...
type SetArea Command

var (
	// SetPoint ...
	SetPoint = func(lat, lon float64) SetArea {
		return SetArea(NewCommand("POINT", floatString(lat), floatString(lon)))
	}

	// SetPointZ ...
	SetPointZ = func(lat, lon, z float64) SetArea {
		return SetArea(NewCommand("POINT", floatString(lat), floatString(lon), floatString(z)))
	}

	// SetFeatureCollection ...
	SetFeatureCollection = func(fc *geojson.FeatureCollection) SetArea {
		b, _ := fc.MarshalJSON()
		return SetArea(NewCommand("OBJECT", string(b)))
	}

	// SetFeature ...
	SetFeature = func(ft *geojson.Feature) SetArea {
		b, _ := ft.MarshalJSON()
		return SetArea(NewCommand("OBJECT", string(b)))
	}

	// SetGeometry ...
	SetGeometry = func(gm *geojson.Geometry) SetArea {
		b, _ := gm.MarshalJSON()
		return SetArea(NewCommand("OBJECT", string(b)))
	}

	// SetHash ...
	SetHash = func(hash string) SetArea {
		return SetArea(NewCommand("HASH", hash))
	}

	// SetString ...
	SetString = func(str string) SetArea {
		return SetArea(NewCommand("STRING", str))
	}
)
