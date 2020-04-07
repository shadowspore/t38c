package t38c

import (
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

// SearchArea ...
type SearchArea Command

// AreaGet ...
func AreaGet(objectID string) SearchArea {
	return SearchArea(NewCommand("GET", objectID))
}

// AreaBounds ...
func AreaBounds(minlat, minlon, maxlat, maxlon float64) SearchArea {
	return SearchArea(NewCommand("BOUNDS", floatString(minlat), floatString(minlon), floatString(maxlat), floatString(maxlon)))
}

// AreaFeatureCollection ...
func AreaFeatureCollection(fc *geojson.FeatureCollection) SearchArea {
	// TODO: handle error?
	b, _ := fc.MarshalJSON()
	return SearchArea(NewCommand("OBJECT", string(b)))
}

// AreaFeature ...
func AreaFeature(ft *geojson.Feature) SearchArea {
	// TODO: handle error?
	b, _ := ft.MarshalJSON()
	return SearchArea(NewCommand("OBJECT", string(b)))
}

// AreaGeometry ...
func AreaGeometry(gm *geojson.Geometry) SearchArea {
	// TODO: handle error?
	b, _ := gm.MarshalJSON()
	return SearchArea(NewCommand("OBJECT", string(b)))

}

// AreaCircle ...
func AreaCircle(lat, lon, meters float64) SearchArea {
	return SearchArea(NewCommand("CIRCLE", floatString(lat), floatString(lon), floatString(meters)))
}

// AreaTile ...
func AreaTile(x, y, z int) SearchArea {
	return SearchArea(NewCommand("TILE", strconv.Itoa(x), strconv.Itoa(y), strconv.Itoa(z)))
}

// AreaQuadkey ...
func AreaQuadkey(quadkey string) SearchArea {
	return SearchArea(NewCommand("QUADKEY", quadkey))
}

// AreaHash ...
func AreaHash(hash string) SearchArea {
	return SearchArea(NewCommand("HASH", hash))
}

// NearbyArea ...
type NearbyArea Command

// NearbyPoint ...
func NearbyPoint(lat, lon, meters float64) NearbyArea {
	return NearbyArea(NewCommand("POINT", floatString(lat), floatString(lon), floatString(meters)))
}

// SetArea ...
type SetArea Command

// SetPoint ...
func SetPoint(lat, lon float64) SetArea {
	return SetArea(NewCommand("POINT", floatString(lat), floatString(lon)))
}

// SetPointZ ...
func SetPointZ(lat, lon, z float64) SetArea {
	return SetArea(NewCommand("POINT", floatString(lat), floatString(lon), floatString(z)))
}

// SetFeatureCollection ...
func SetFeatureCollection(fc *geojson.FeatureCollection) SetArea {
	b, _ := fc.MarshalJSON()
	return SetArea(NewCommand("OBJECT", string(b)))
}

// SetFeature ...
func SetFeature(ft *geojson.Feature) SetArea {
	b, _ := ft.MarshalJSON()
	return SetArea(NewCommand("OBJECT", string(b)))
}

// SetGeometry ...
func SetGeometry(gm *geojson.Geometry) SetArea {
	b, _ := gm.MarshalJSON()
	return SetArea(NewCommand("OBJECT", string(b)))
}

// SetHash ...
func SetHash(hash string) SetArea {
	return SetArea(NewCommand("HASH", hash))
}

// SetString ...
func SetString(str string) SetArea {
	return SetArea(NewCommand("STRING", str))
}
