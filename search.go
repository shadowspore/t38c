package t38c

import (
	"encoding/json"
)

func (client *Tile38Client) searchObjects(cmd, key string, area Command, opts []SearchOption) ([]*GeoJSONObject, error) {
	var resp struct {
		Fields  []string `json:"fields"`
		Objects []struct {
			ID     string          `json:"id"`
			Object json.RawMessage `json:"object"`
			Fields []float64       `json:"fields"`
		} `json:"objects"`
	}

	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	objects := make([]*GeoJSONObject, len(resp.Objects))
	haveFields := len(resp.Fields) > 0
	for idx, obj := range resp.Objects {
		geoObj := &GeoJSONObject{}
		geoObj.Tile38ID = obj.ID
		geo, err := unmarshalGeoJSON(obj.Object)
		if err != nil {
			return nil, err
		}

		geoObj.GeoJSON = geo
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

func (client *Tile38Client) searchPoints(cmd, key string, area Command, opts []SearchOption) ([]*PointObject, error) {
	var resp struct {
		Fields []string `json:"fields"`
		Points []struct {
			ID     string    `json:"id"`
			Point  Point     `json:"point"`
			Fields []float64 `json:"fields"`
		} `json:"points"`
	}

	opts = append(opts, SearchOption(NewCommand("POINTS")))
	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
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

func (client *Tile38Client) searchIDs(cmd, key string, area Command, opts []SearchOption) ([]string, error) {
	var resp struct {
		IDs []string `json:"ids"`
	}

	opts = append(opts, SearchOption(NewCommand("IDS")))
	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return resp.IDs, nil
}

func (client *Tile38Client) searchBounds(cmd, key string, area Command, opts []SearchOption) ([]*BoundsObject, error) {
	var resp struct {
		Fields []string `json:"fields"`
		Bounds []struct {
			ID     string    `json:"id"`
			Bounds Bounds    `json:"bounds"`
			Fields []float64 `json:"fields"`
		} `json:"bounds"`
	}

	opts = append(opts, SearchOption(NewCommand("BOUNDS")))
	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	objects := make([]*BoundsObject, len(resp.Bounds))
	haveFields := len(resp.Fields) > 0
	for idx, obj := range resp.Bounds {
		boundObj := &BoundsObject{}
		boundObj.Tile38ID = obj.ID
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

func (client *Tile38Client) searchHashes(cmd, key string, area Command, precision int, opts []SearchOption) ([]*HashObject, error) {
	var resp struct {
		Fields []string `json:"fields"`
		Hashes []struct {
			ID     string    `json:"id"`
			Hash   string    `json:"hash"`
			Fields []float64 `json:"fields"`
		} `json:"hashes"`
	}

	opts = append(opts, SearchOption(NewCommand("HASHES", precision)))
	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	objects := make([]*HashObject, len(resp.Hashes))
	haveFields := len(resp.Fields) > 0
	for idx, obj := range resp.Hashes {
		hashObj := &HashObject{}
		hashObj.Tile38ID = obj.ID
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

func buildArgs(key string, area Command, opts []SearchOption) []interface{} {
	var args []interface{}
	args = append(args, key)
	for _, opt := range opts {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}
	args = append(args, area.Name)
	args = append(args, area.Args...)
	return args
}

// Intersects searches a collection for objects that intersect a specified bounding area.
func (client *Tile38Client) Intersects(key string, area SearchArea, opts ...SearchOption) ([]*GeoJSONObject, error) {
	return client.searchObjects("INTERSECTS", key, Command(area), opts)
}

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func (client *Tile38Client) Within(key string, area SearchArea, opts ...SearchOption) ([]*GeoJSONObject, error) {
	return client.searchObjects("WITHIN", key, Command(area), opts)
}

// Nearby searches a collection for objects that are close to a specified point.
func (client *Tile38Client) Nearby(key string, area NearbyArea, opts ...SearchOption) ([]*GeoJSONObject, error) {
	return client.searchObjects("NEARBY", key, Command(area), opts)
}

// IntersectsPoints ...
func (client *Tile38Client) IntersectsPoints(key string, area SearchArea, opts ...SearchOption) ([]*PointObject, error) {
	return client.searchPoints("INTERSECTS", key, Command(area), opts)
}

// WithinPoints ...
func (client *Tile38Client) WithinPoints(key string, area SearchArea, opts ...SearchOption) ([]*PointObject, error) {
	return client.searchPoints("WITHIN", key, Command(area), opts)
}

// NearbyPoints ...
func (client *Tile38Client) NearbyPoints(key string, area NearbyArea, opts ...SearchOption) ([]*PointObject, error) {
	return client.searchPoints("NEARBY", key, Command(area), opts)
}

// IntersectsIDs ...
func (client *Tile38Client) IntersectsIDs(key string, area SearchArea, opts ...SearchOption) ([]string, error) {
	return client.searchIDs("INTERSECTS", key, Command(area), opts)
}

// WithinIDs ...
func (client *Tile38Client) WithinIDs(key string, area SearchArea, opts ...SearchOption) ([]string, error) {
	return client.searchIDs("WITHIN", key, Command(area), opts)
}

// NearbyIDs ...
func (client *Tile38Client) NearbyIDs(key string, area NearbyArea, opts ...SearchOption) ([]string, error) {
	return client.searchIDs("NEARBY", key, Command(area), opts)
}

// IntersectsBounds ...
func (client *Tile38Client) IntersectsBounds(key string, area SearchArea, opts ...SearchOption) ([]*BoundsObject, error) {
	return client.searchBounds("INTERSECTS", key, Command(area), opts)
}

// WithinBounds ...
func (client *Tile38Client) WithinBounds(key string, area SearchArea, opts ...SearchOption) ([]*BoundsObject, error) {
	return client.searchBounds("WITHIN", key, Command(area), opts)
}

// NearbyBounds ...
func (client *Tile38Client) NearbyBounds(key string, area NearbyArea, opts ...SearchOption) ([]*BoundsObject, error) {
	return client.searchBounds("NEARBY", key, Command(area), opts)
}

// IntersectsHashes ...
func (client *Tile38Client) IntersectsHashes(key string, area SearchArea, precision int, opts ...SearchOption) ([]*HashObject, error) {
	return client.searchHashes("INTERSECTS", key, Command(area), precision, opts)
}

// WithinHashes ...
func (client *Tile38Client) WithinHashes(key string, area SearchArea, precision int, opts ...SearchOption) ([]*HashObject, error) {
	return client.searchHashes("WITHIN", key, Command(area), precision, opts)
}

// NearbyHashes ...
func (client *Tile38Client) NearbyHashes(key string, area NearbyArea, precision int, opts ...SearchOption) ([]*HashObject, error) {
	return client.searchHashes("NEARBY", key, Command(area), precision, opts)
}
