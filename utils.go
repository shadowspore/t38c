package t38c

import (
	"strconv"

	geojson "github.com/paulmach/go.geojson"
	"github.com/tidwall/gjson"
)

func floatToString(val float64) string {
	return strconv.FormatFloat(val, 'f', 6, 64)
}

func unmarshalGeoJSON(data []byte) (interface{}, error) {
	switch gjson.GetBytes(data, "type").String() {
	case "FeatureCollection":
		return geojson.UnmarshalFeatureCollection(data)
	case "Feature":
		return geojson.UnmarshalFeature(data)
	default:
		return geojson.UnmarshalGeometry(data)
	}
}
