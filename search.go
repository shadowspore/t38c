package t38c

import (
	"encoding/json"
	"strconv"
)

type searchResponse struct {
	Cursor int `json:"cursor"`
	Count  int `json:"count"`
}

// SearchObjectsResponse struct
type SearchObjectsResponse struct {
	searchResponse
	Objects []*GeoJSONObject
}

func (client *Tile38Client) searchObjects(cmd, key string, area Command, opts []SearchOption) (*SearchObjectsResponse, error) {
	var resp struct {
		searchResponse
		Fields  []string `json:"fields"`
		Objects []struct {
			ID       string          `json:"id"`
			Object   json.RawMessage `json:"object"`
			Fields   []float64       `json:"fields"`
			Distance *float64        `json:"distance,omitempty"`
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
		geoObj.Distance = obj.Distance
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

	return &SearchObjectsResponse{
		searchResponse: resp.searchResponse,
		Objects:        objects,
	}, nil
}

// SearchPointsResponse struct
type SearchPointsResponse struct {
	searchResponse
	Points []*PointObject
}

func (client *Tile38Client) searchPoints(cmd, key string, area Command, opts []SearchOption) (*SearchPointsResponse, error) {
	var resp struct {
		searchResponse
		Fields []string `json:"fields"`
		Points []struct {
			ID       string    `json:"id"`
			Point    Point     `json:"point"`
			Fields   []float64 `json:"fields"`
			Distance *float64  `json:"distance,omitempty"`
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

	return &SearchPointsResponse{
		searchResponse: resp.searchResponse,
		Points:         points,
	}, nil
}

// SearchIDsResponse struct
type SearchIDsResponse struct {
	searchResponse
	IDs []string
}

func (client *Tile38Client) searchIDs(cmd, key string, area Command, opts []SearchOption) (*SearchIDsResponse, error) {
	var resp struct {
		searchResponse
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

	return &SearchIDsResponse{
		searchResponse: resp.searchResponse,
		IDs:            resp.IDs,
	}, nil
}

// SearchBoundsResponse struct
type SearchBoundsResponse struct {
	searchResponse
	Bounds []*BoundsObject
}

func (client *Tile38Client) searchBounds(cmd, key string, area Command, opts []SearchOption) (*SearchBoundsResponse, error) {
	var resp struct {
		searchResponse
		Fields []string `json:"fields"`
		Bounds []struct {
			ID       string    `json:"id"`
			Bounds   Bounds    `json:"bounds"`
			Fields   []float64 `json:"fields"`
			Distance *float64  `json:"distance,omitempty"`
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

	return &SearchBoundsResponse{
		searchResponse: resp.searchResponse,
		Bounds:         objects,
	}, nil
}

// SearchHashesResponse struct
type SearchHashesResponse struct {
	searchResponse
	Hashes []*HashObject
}

func (client *Tile38Client) searchHashes(cmd, key string, area Command, precision int, opts []SearchOption) (*SearchHashesResponse, error) {
	var resp struct {
		searchResponse
		Fields []string `json:"fields"`
		Hashes []struct {
			ID       string    `json:"id"`
			Hash     string    `json:"hash"`
			Fields   []float64 `json:"fields"`
			Distance *float64  `json:"distance,omitempty"`
		} `json:"hashes"`
	}

	opts = append(opts, SearchOption(NewCommand("HASHES", strconv.Itoa(precision))))
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

	return &SearchHashesResponse{
		searchResponse: resp.searchResponse,
		Hashes:         objects,
	}, nil
}

// SearchCountResponse struct
type SearchCountResponse struct {
	searchResponse
}

func (client *Tile38Client) searchCount(cmd, key string, area Command, opts []SearchOption) (*SearchCountResponse, error) {
	var resp struct {
		searchResponse
	}

	opts = append(opts, SearchOption(NewCommand("COUNT")))
	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return &SearchCountResponse{}, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return &SearchCountResponse{}, err
	}

	return &SearchCountResponse{
		searchResponse: resp.searchResponse,
	}, nil
}

func buildArgs(key string, area Command, opts []SearchOption) []string {
	var args []string
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
func (client *Tile38Client) Intersects(key string, area SearchArea, opts ...SearchOption) (*SearchObjectsResponse, error) {
	return client.searchObjects("INTERSECTS", key, Command(area), opts)
}

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func (client *Tile38Client) Within(key string, area SearchArea, opts ...SearchOption) (*SearchObjectsResponse, error) {
	return client.searchObjects("WITHIN", key, Command(area), opts)
}

// Nearby searches a collection for objects that are close to a specified point.
func (client *Tile38Client) Nearby(key string, area NearbyArea, opts ...SearchOption) (*SearchObjectsResponse, error) {
	return client.searchObjects("NEARBY", key, Command(area), opts)
}

// IntersectsPoints ...
func (client *Tile38Client) IntersectsPoints(key string, area SearchArea, opts ...SearchOption) (*SearchPointsResponse, error) {
	return client.searchPoints("INTERSECTS", key, Command(area), opts)
}

// WithinPoints ...
func (client *Tile38Client) WithinPoints(key string, area SearchArea, opts ...SearchOption) (*SearchPointsResponse, error) {
	return client.searchPoints("WITHIN", key, Command(area), opts)
}

// NearbyPoints ...
func (client *Tile38Client) NearbyPoints(key string, area NearbyArea, opts ...SearchOption) (*SearchPointsResponse, error) {
	return client.searchPoints("NEARBY", key, Command(area), opts)
}

// IntersectsIDs ...
func (client *Tile38Client) IntersectsIDs(key string, area SearchArea, opts ...SearchOption) (*SearchIDsResponse, error) {
	return client.searchIDs("INTERSECTS", key, Command(area), opts)
}

// WithinIDs ...
func (client *Tile38Client) WithinIDs(key string, area SearchArea, opts ...SearchOption) (*SearchIDsResponse, error) {
	return client.searchIDs("WITHIN", key, Command(area), opts)
}

// NearbyIDs ...
func (client *Tile38Client) NearbyIDs(key string, area NearbyArea, opts ...SearchOption) (*SearchIDsResponse, error) {
	return client.searchIDs("NEARBY", key, Command(area), opts)
}

// IntersectsBounds ...
func (client *Tile38Client) IntersectsBounds(key string, area SearchArea, opts ...SearchOption) (*SearchBoundsResponse, error) {
	return client.searchBounds("INTERSECTS", key, Command(area), opts)
}

// WithinBounds ...
func (client *Tile38Client) WithinBounds(key string, area SearchArea, opts ...SearchOption) (*SearchBoundsResponse, error) {
	return client.searchBounds("WITHIN", key, Command(area), opts)
}

// NearbyBounds ...
func (client *Tile38Client) NearbyBounds(key string, area NearbyArea, opts ...SearchOption) (*SearchBoundsResponse, error) {
	return client.searchBounds("NEARBY", key, Command(area), opts)
}

// IntersectsHashes ...
func (client *Tile38Client) IntersectsHashes(key string, area SearchArea, precision int, opts ...SearchOption) (*SearchHashesResponse, error) {
	return client.searchHashes("INTERSECTS", key, Command(area), precision, opts)
}

// WithinHashes ...
func (client *Tile38Client) WithinHashes(key string, area SearchArea, precision int, opts ...SearchOption) (*SearchHashesResponse, error) {
	return client.searchHashes("WITHIN", key, Command(area), precision, opts)
}

// NearbyHashes ...
func (client *Tile38Client) NearbyHashes(key string, area NearbyArea, precision int, opts ...SearchOption) (*SearchHashesResponse, error) {
	return client.searchHashes("NEARBY", key, Command(area), precision, opts)
}

// IntersectsCount ...
func (client *Tile38Client) IntersectsCount(key string, area SearchArea, opts ...SearchOption) (*SearchCountResponse, error) {
	return client.searchCount("INTERSECTS", key, Command(area), opts)
}

// WithinCount ...
func (client *Tile38Client) WithinCount(key string, area SearchArea, opts ...SearchOption) (*SearchCountResponse, error) {
	return client.searchCount("WITHIN", key, Command(area), opts)
}

// NearbyCount ...
func (client *Tile38Client) NearbyCount(key string, area NearbyArea, opts ...SearchOption) (*SearchCountResponse, error) {
	return client.searchCount("NEARBY", key, Command(area), opts)
}
