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

func unmarshalObjects(data []byte) ([]*Object, error) {
	var resp struct {
		Fields  []string `json:"fields"`
		Objects []struct {
			ID       string          `json:"id"`
			Object   json.RawMessage `json:"object"`
			Fields   []float64       `json:"fields"`
			Distance *float64        `json:"distance,omitempty"`
		} `json:"objects"`
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	objects := make([]*Object, len(resp.Objects))
	haveFields := len(resp.Fields) > 0
	for idx, obj := range resp.Objects {
		geoObj := &Object{}
		geoObj.Tile38ID = obj.ID
		geoObj.Distance = obj.Distance
		ob, err := unmarshalObject(obj.Object)
		if err != nil {
			return nil, err
		}

		geoObj.Object = ob
		if haveFields {
			geoObj.Fields = make(map[string]float64)
			for fieldIdx, field := range resp.Fields {
				geoObj.Fields[field] = obj.Fields[fieldIdx]
			}
		}

		objects[idx] = geoObj
	}

	return objects, nil
}

func unmarshalPoints(data []byte) ([]*PointObject, error) {
	var resp struct {
		Fields []string `json:"fields"`
		Points []struct {
			ID       string    `json:"id"`
			Point    Point     `json:"point"`
			Fields   []float64 `json:"fields"`
			Distance *float64  `json:"distance,omitempty"`
		} `json:"points"`
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	points := make([]*PointObject, len(resp.Points))
	haveFields := len(resp.Fields) > 0
	for idx, point := range resp.Points {
		pointObj := &PointObject{}
		pointObj.Tile38ID = point.ID
		pointObj.Distance = point.Distance
		pointObj.Point = point.Point
		if haveFields {
			pointObj.Fields = make(map[string]float64)
			for fieldIdx, field := range resp.Fields {
				pointObj.Fields[field] = point.Fields[fieldIdx]
			}
		}

		points[idx] = pointObj
	}

	return points, nil
}

func unmarshalBounds(data []byte) ([]*BoundsObject, error) {
	var resp struct {
		Fields []string `json:"fields"`
		Bounds []struct {
			ID       string    `json:"id"`
			Bounds   Bounds    `json:"bounds"`
			Fields   []float64 `json:"fields"`
			Distance *float64  `json:"distance,omitempty"`
		} `json:"bounds"`
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	objects := make([]*BoundsObject, len(resp.Bounds))
	haveFields := len(resp.Fields) > 0
	for idx, obj := range resp.Bounds {
		boundObj := &BoundsObject{}
		boundObj.Tile38ID = obj.ID
		boundObj.Distance = obj.Distance
		boundObj.Bounds = obj.Bounds
		if haveFields {
			boundObj.Fields = make(map[string]float64)
			for fieldIdx, field := range resp.Fields {
				boundObj.Fields[field] = obj.Fields[fieldIdx]
			}
		}

		objects[idx] = boundObj
	}

	return objects, nil
}

func unmarshalIDs(data []byte) ([]string, error) {
	var resp struct {
		IDs []string `json:"ids"`
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	return resp.IDs, nil
}

func unmarshalHashes(data []byte) ([]*HashObject, error) {
	var resp struct {
		Fields []string `json:"fields"`
		Hashes []struct {
			ID       string    `json:"id"`
			Hash     string    `json:"hash"`
			Fields   []float64 `json:"fields"`
			Distance *float64  `json:"distance,omitempty"`
		} `json:"hashes"`
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}

	objects := make([]*HashObject, len(resp.Hashes))
	haveFields := len(resp.Fields) > 0
	for idx, obj := range resp.Hashes {
		hashObj := &HashObject{}
		hashObj.Tile38ID = obj.ID
		hashObj.Distance = obj.Distance
		hashObj.Hash = obj.Hash
		if haveFields {
			hashObj.Fields = make(map[string]float64)
			for fieldIdx, field := range resp.Fields {
				hashObj.Fields[field] = obj.Fields[fieldIdx]
			}
		}

		objects[idx] = hashObj
	}

	return objects, nil
}
