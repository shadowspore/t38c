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
	Fields  []string `json:"fields,omitempty"`
	Objects []struct {
		ID       string    `json:"ID"`
		Object   Object    `json:"object"`
		Fields   []float64 `json:"fields"`
		Distance *float64  `json:"distance"`
	} `json:"objects"`
}

func (client *Tile38Client) searchObjects(cmd, key string, area Command, opts []SearchOption) (*SearchObjectsResponse, error) {
	var resp *SearchObjectsResponse

	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &resp)
	return resp, nil
}

// SearchPointsResponse struct
type SearchPointsResponse struct {
	searchResponse
	Fields []string `json:"fields,omitempty"`
	Points []struct {
		ID       string    `json:"ID"`
		Point    Point     `json:"point"`
		Fields   []float64 `json:"fields,omitempty"`
		Distance *float64  `json:"distance,omitempty"`
	} `json:"points"`
}

func (client *Tile38Client) searchPoints(cmd, key string, area Command, opts []SearchOption) (*SearchPointsResponse, error) {
	var resp *SearchPointsResponse

	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &resp)
	return resp, nil
}

// SearchIDsResponse struct
type SearchIDsResponse struct {
	searchResponse
	IDs []string `json:"ids"`
}

func (client *Tile38Client) searchIDs(cmd, key string, area Command, opts []SearchOption) (*SearchIDsResponse, error) {
	var resp *SearchIDsResponse

	opts = append(opts, SearchOption(NewCommand("IDS")))
	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &resp)
	return resp, err
}

// SearchBoundsResponse struct
type SearchBoundsResponse struct {
	searchResponse
	Fields []string `json:"fields,omitempty"`
	Bounds []struct {
		ID       string    `json:"ID"`
		Bounds   Bounds    `json:"bounds"`
		Fields   []float64 `json:"fields,omitempty"`
		Distance *float64  `json:"distance,omitempty"`
	} `json:"bounds"`
}

func (client *Tile38Client) searchBounds(cmd, key string, area Command, opts []SearchOption) (*SearchBoundsResponse, error) {
	var resp *SearchBoundsResponse

	opts = append(opts, SearchOption(NewCommand("BOUNDS")))
	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &resp)
	return resp, err
}

// SearchHashesResponse struct
type SearchHashesResponse struct {
	searchResponse
	Fields []string `json:"fields,omitempty"`
	Hashes []struct {
		ID       string    `json:"id"`
		Hash     string    `json:"hash"`
		Fields   []float64 `json:"fields,omitempty"`
		Distance *float64  `json:"distance,omitempty"`
	} `json:"hashes"`
}

func (client *Tile38Client) searchHashes(cmd, key string, area Command, precision int, opts []SearchOption) (*SearchHashesResponse, error) {
	var resp *SearchHashesResponse

	opts = append(opts, SearchOption(NewCommand("HASHES", strconv.Itoa(precision))))
	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &resp)
	return resp, err
}

// SearchCountResponse struct
type SearchCountResponse struct {
	searchResponse
}

func (client *Tile38Client) searchCount(cmd, key string, area Command, opts []SearchOption) (*SearchCountResponse, error) {
	var resp *SearchCountResponse

	opts = append(opts, SearchOption(NewCommand("COUNT")))
	args := buildArgs(key, area, opts)
	b, err := client.Execute(cmd, args...)
	if err != nil {
		return &SearchCountResponse{}, err
	}

	err = json.Unmarshal(b, &resp)
	return resp, err
}

func buildArgs(key string, area Command, opts []SearchOption) []string {
	var args []string
	args = append(args, key)
	for _, opt := range opts {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}

	if len(area.Name) > 0 {
		args = append(args, area.Name)
		args = append(args, area.Args...)
	}
	return args
}

// Intersects searches a collection for objects that intersect a specified bounding area.
func (client *Tile38Client) Intersects(key string, area SearchArea, opts ...SearchOption) (*SearchObjectsResponse, error) {
	return client.searchObjects("INTERSECTS", key, Command(area), opts)
}

// IntersectsPoints ...
func (client *Tile38Client) IntersectsPoints(key string, area SearchArea, opts ...SearchOption) (*SearchPointsResponse, error) {
	return client.searchPoints("INTERSECTS", key, Command(area), opts)
}

// IntersectsIDs ...
func (client *Tile38Client) IntersectsIDs(key string, area SearchArea, opts ...SearchOption) (*SearchIDsResponse, error) {
	return client.searchIDs("INTERSECTS", key, Command(area), opts)
}

// IntersectsBounds ...
func (client *Tile38Client) IntersectsBounds(key string, area SearchArea, opts ...SearchOption) (*SearchBoundsResponse, error) {
	return client.searchBounds("INTERSECTS", key, Command(area), opts)
}

// IntersectsHashes ...
func (client *Tile38Client) IntersectsHashes(key string, area SearchArea, precision int, opts ...SearchOption) (*SearchHashesResponse, error) {
	return client.searchHashes("INTERSECTS", key, Command(area), precision, opts)
}

// IntersectsCount ...
func (client *Tile38Client) IntersectsCount(key string, area SearchArea, opts ...SearchOption) (*SearchCountResponse, error) {
	return client.searchCount("INTERSECTS", key, Command(area), opts)
}

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func (client *Tile38Client) Within(key string, area SearchArea, opts ...SearchOption) (*SearchObjectsResponse, error) {
	return client.searchObjects("WITHIN", key, Command(area), opts)
}

// WithinPoints ...
func (client *Tile38Client) WithinPoints(key string, area SearchArea, opts ...SearchOption) (*SearchPointsResponse, error) {
	return client.searchPoints("WITHIN", key, Command(area), opts)
}

// WithinIDs ...
func (client *Tile38Client) WithinIDs(key string, area SearchArea, opts ...SearchOption) (*SearchIDsResponse, error) {
	return client.searchIDs("WITHIN", key, Command(area), opts)
}

// WithinBounds ...
func (client *Tile38Client) WithinBounds(key string, area SearchArea, opts ...SearchOption) (*SearchBoundsResponse, error) {
	return client.searchBounds("WITHIN", key, Command(area), opts)
}

// WithinHashes ...
func (client *Tile38Client) WithinHashes(key string, area SearchArea, precision int, opts ...SearchOption) (*SearchHashesResponse, error) {
	return client.searchHashes("WITHIN", key, Command(area), precision, opts)
}

// WithinCount ...
func (client *Tile38Client) WithinCount(key string, area SearchArea, opts ...SearchOption) (*SearchCountResponse, error) {
	return client.searchCount("WITHIN", key, Command(area), opts)
}

// Nearby searches a collection for objects that are close to a specified point.
func (client *Tile38Client) Nearby(key string, area NearbyArea, opts ...SearchOption) (*SearchObjectsResponse, error) {
	return client.searchObjects("NEARBY", key, Command(area), opts)
}

// NearbyPoints ...
func (client *Tile38Client) NearbyPoints(key string, area NearbyArea, opts ...SearchOption) (*SearchPointsResponse, error) {
	return client.searchPoints("NEARBY", key, Command(area), opts)
}

// NearbyIDs ...
func (client *Tile38Client) NearbyIDs(key string, area NearbyArea, opts ...SearchOption) (*SearchIDsResponse, error) {
	return client.searchIDs("NEARBY", key, Command(area), opts)
}

// NearbyBounds ...
func (client *Tile38Client) NearbyBounds(key string, area NearbyArea, opts ...SearchOption) (*SearchBoundsResponse, error) {
	return client.searchBounds("NEARBY", key, Command(area), opts)
}

// NearbyHashes ...
func (client *Tile38Client) NearbyHashes(key string, area NearbyArea, precision int, opts ...SearchOption) (*SearchHashesResponse, error) {
	return client.searchHashes("NEARBY", key, Command(area), precision, opts)
}

// NearbyCount ...
func (client *Tile38Client) NearbyCount(key string, area NearbyArea, opts ...SearchOption) (*SearchCountResponse, error) {
	return client.searchCount("NEARBY", key, Command(area), opts)
}

// Scan ...
func (client *Tile38Client) Scan(key string, opts ...SearchOption) (*SearchObjectsResponse, error) {
	return client.searchObjects("SCAN", key, NewCommand(""), opts)
}

// ScanPoints ...
func (client *Tile38Client) ScanPoints(key string, opts ...SearchOption) (*SearchPointsResponse, error) {
	return client.searchPoints("SCAN", key, NewCommand(""), opts)
}

// ScanIDs ...
func (client *Tile38Client) ScanIDs(key string, opts ...SearchOption) (*SearchIDsResponse, error) {
	return client.searchIDs("SCAN", key, NewCommand(""), opts)
}

// ScanBounds ...
func (client *Tile38Client) ScanBounds(key string, opts ...SearchOption) (*SearchBoundsResponse, error) {
	return client.searchBounds("SCAN", key, NewCommand(""), opts)
}

// ScanHashes ...
func (client *Tile38Client) ScanHashes(key string, precision int, opts ...SearchOption) (*SearchHashesResponse, error) {
	return client.searchHashes("SCAN", key, NewCommand(""), precision, opts)
}

// Search ...
func (client *Tile38Client) Search(key string, opts ...SearchOption) (*SearchObjectsResponse, error) {
	return client.searchObjects("SEARCH", key, NewCommand(""), opts)
}

// SearchCount ...
func (client *Tile38Client) SearchCount(key string, opts ...SearchOption) (*SearchCountResponse, error) {
	return client.searchCount("SEARCH", key, NewCommand(""), opts)
}

// SearchIDs ...
func (client *Tile38Client) SearchIDs(key string, opts ...SearchOption) (*SearchIDsResponse, error) {
	return client.searchIDs("SEARCH", key, NewCommand(""), opts)
}
