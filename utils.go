package t38c

import (
	"encoding/json"
	"strconv"

	geojson "github.com/paulmach/go.geojson"
	"github.com/tidwall/gjson"
)

func unmarshalObject(data []byte) (interface{}, error) {
	tp := gjson.GetBytes(data, "type")
	if !tp.Exists() {
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return nil, err
		}
		return str, nil
	}

	switch tp.String() {
	case "FeatureCollection":
		return geojson.UnmarshalFeatureCollection(data)
	case "Feature":
		return geojson.UnmarshalFeature(data)
	default:
		return geojson.UnmarshalGeometry(data)
	}
}

func floatString(val float64) string {
	return strconv.FormatFloat(val, 'f', 10, 64)
}
