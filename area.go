package t38c

import (
	geojson "github.com/paulmach/go.geojson"
)

// SearchArea ...
type SearchArea string

// AreaGet ...
func AreaGet(objectID string) SearchArea {
	return SearchArea("GET " + objectID)
}

// AreaBounds ...
func AreaBounds(minlat, minlon, maxlat, maxlon float64) SearchArea {
	return SearchArea(
		"BOUNDS " +
			floatToString(minlat) + " " +
			floatToString(minlon) + " " +
			floatToString(maxlat) + " " +
			floatToString(maxlon),
	)
}

// AreaFeatureCollection ...
func AreaFeatureCollection(fc *geojson.FeatureCollection) SearchArea {
	// TODO: handle error?
	b, _ := fc.MarshalJSON()
	return SearchArea(
		"OBJECT " + string(b),
	)
}

// AreaFeature ...
func AreaFeature(ft *geojson.Feature) SearchArea {
	// TODO: handle error?
	b, _ := ft.MarshalJSON()
	return SearchArea(
		"OBJECT " + string(b),
	)
}

// AreaGeometry ...
func AreaGeometry(gm *geojson.Geometry) SearchArea {
	// TODO: handle error?
	b, _ := gm.MarshalJSON()
	return SearchArea(
		"OBJECT " + string(b),
	)

}

// AreaCircle ...
func AreaCircle(lat, lon, meters float64) SearchArea {
	return SearchArea(
		"CIRCLE " +
			floatToString(lat) + " " +
			floatToString(lon) + " " +
			floatToString(meters),
	)
}

// NearbyArea ...
type NearbyArea string

// NearbyPoint ...
func NearbyPoint(lat, lon, meters float64) NearbyArea {
	return NearbyArea(
		"POINT " + floatToString(lat) + " " + floatToString(lon) + " " + floatToString(meters),
	)
}

// SetArea ...
type SetArea string

// SetPoint ...
func SetPoint(lat, lon float64) SetArea {
	return SetArea(
		"POINT " + floatToString(lat) + " " + floatToString(lon),
	)
}

// SetPointZ ...
func SetPointZ(lat, lon, z float64) SetArea {
	return SetArea(
		"POINT " + floatToString(lat) + " " + floatToString(lon) + " " + floatToString(z),
	)
}

// SetFeatureCollection ...
func SetFeatureCollection(fc *geojson.FeatureCollection) SetArea {
	b, _ := fc.MarshalJSON()
	return SetArea(
		"OBJECT " + string(b),
	)
}

// SetFeature ...
func SetFeature(ft *geojson.Feature) SetArea {
	b, _ := ft.MarshalJSON()
	return SetArea(
		"OBJECT " + string(b),
	)
}

// SetGeometry ...
func SetGeometry(gm *geojson.Geometry) SetArea {
	b, _ := gm.MarshalJSON()
	return SetArea(
		"OBJECT " + string(b),
	)
}
