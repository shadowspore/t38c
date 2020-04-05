package t38c

import (
	"strings"
)

func (client *Tile38Client) searchObjects(cmd, key, area string, opts []SearchOption) ([]*GeoJSONObject, error) {
	var resp struct {
		Fields  []string `json:"fields"`
		Objects []struct {
			ID     string         `json:"id"`
			Object *GeoJSONObject `json:"object"`
			Fields []float64      `json:"fields"`
		} `json:"objects"`
	}

	command := buildSearchCommand(cmd, key, area, opts)
	if err := client.execute(command, &resp); err != nil {
		return nil, err
	}

	objects := make([]*GeoJSONObject, len(resp.Objects))
	haveFields := len(resp.Fields) > 0
	for idx, obj := range resp.Objects {
		o := obj.Object
		o.Tile38ID = obj.ID
		if haveFields {
			o.Fields = make(map[string]float64)
			for fieldIdx, field := range resp.Fields {
				o.Fields[field] = obj.Fields[fieldIdx]
			}
		}

		objects[idx] = o
	}

	return objects, nil
}

func buildSearchCommand(cmd, key, area string, opts []SearchOption) string {
	var sb strings.Builder
	sb.WriteString(cmd + " " + key)
	for _, opt := range opts {
		sb.WriteString(" " + string(opt))
	}
	sb.WriteString(" " + area)
	return sb.String()
}

// Intersects searches a collection for objects that intersect a specified bounding area.
func (client *Tile38Client) Intersects(key string, area SearchArea, opts ...SearchOption) ([]*GeoJSONObject, error) {
	return client.searchObjects("INTERSECTS", key, string(area), opts)
}

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func (client *Tile38Client) Within(key string, area SearchArea, opts ...SearchOption) ([]*GeoJSONObject, error) {
	return client.searchObjects("WITHIN", key, string(area), opts)
}

// Nearby searches a collection for objects that are close to a specified point.
func (client *Tile38Client) Nearby(key string, area NearbyArea, opts ...SearchOption) ([]*GeoJSONObject, error) {
	return client.searchObjects("NEARBY", key, string(area), opts)
}

func (client *Tile38Client) searchPoints(cmd, key, area string, opts []SearchOption) ([]*PointObject, error) {
	var resp struct {
		Fields []string `json:"fields"`
		Points []struct {
			ID     string    `json:"id"`
			Point  Point     `json:"point"`
			Fields []float64 `json:"fields"`
		} `json:"points"`
	}

	opts = append(opts, SearchOption("POINTS"))
	command := buildSearchCommand(cmd, key, area, opts)
	if err := client.execute(command, &resp); err != nil {
		return nil, err
	}

	points := make([]*PointObject, len(resp.Points))
	haveFields := len(resp.Fields) > 0
	for idx, point := range resp.Points {
		pointObj := &PointObject{}
		pointObj.Tile38ID = point.ID
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

// IntersectsPoints ...
func (client *Tile38Client) IntersectsPoints(key string, area SearchArea, opts ...SearchOption) ([]*PointObject, error) {
	return client.searchPoints("INTERSECTS", key, string(area), opts)
}

// WithinPoints ...
func (client *Tile38Client) WithinPoints(key string, area SearchArea, opts ...SearchOption) ([]*PointObject, error) {
	return client.searchPoints("WITHIN", key, string(area), opts)
}

// NearbyPoints ...
func (client *Tile38Client) NearbyPoints(key string, area NearbyArea, opts ...SearchOption) ([]*PointObject, error) {
	return client.searchPoints("NEARBY", key, string(area), opts)
}