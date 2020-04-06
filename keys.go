package t38c

import (
	"encoding/json"
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

// Bounds returns the minimum bounding rectangle for all objects in a key.
func (client *Tile38Client) Bounds(key string) ([][][]float64, error) {
	var resp struct {
		Bounds geojson.Geometry `json:"bounds"`
	}

	b, err := client.Execute("BOUNDS", key)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return resp.Bounds.Polygon, nil
}

// Del remove a specified object.
func (client *Tile38Client) Del(key, objectID string) error {
	_, err := client.Execute("DEL", key)
	return err
}

// Drop remove all objects from specified key.
func (client *Tile38Client) Drop(key string) error {
	_, err := client.Execute("DROP", key)
	return err
}

// Set the value of an id. If a value is already associated to that key/id, it’ll be overwritten.
func (client *Tile38Client) Set(key, objectID string, area SetArea, opts ...SetOption) error {
	var args []string
	args = append(args, key)
	args = append(args, objectID)
	for _, opt := range opts {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}
	args = append(args, area.Name)
	args = append(args, area.Args...)

	_, err := client.Execute("SET", args...)
	return err
}

// Keys returns all keys matching pattern.
func (client *Tile38Client) Keys(pattern string) ([]string, error) {
	var resp struct {
		Keys []string `json:"keys"`
	}

	b, err := client.Execute("KEYS", pattern)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return resp.Keys, nil
}

// Get GeoJSON object.
func (client *Tile38Client) Get(key, objectID string, withFields bool) (*GeoJSONObject, error) {
	return client.GetObject(key, objectID, withFields)
}

// GetObject returns GeoJSON object of an id.
func (client *Tile38Client) GetObject(key, objectID string, withFields bool) (*GeoJSONObject, error) {
	var resp struct {
		Object json.RawMessage    `json:"object"`
		Fields map[string]float64 `json:"fields"`
	}

	var args []string
	args = append(args, key)
	args = append(args, objectID)
	if withFields {
		args = append(args, "WITHFIELDS")
	}

	b, err := client.Execute("GET", args...)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	geo, err := unmarshalGeoJSON(resp.Object)
	if err != nil {
		return nil, err
	}

	obj := &GeoJSONObject{
		Object: Object{
			Tile38ID: key,
			Fields:   resp.Fields,
		},
		GeoJSON: geo,
	}
	return obj, nil
}

// GetPoint get latitude, longitude point.
func (client *Tile38Client) GetPoint(key, objectID string) (Point, error) {
	var resp struct {
		Point Point `json:"point"`
	}

	b, err := client.Execute("GET", key, objectID, "POINT")
	if err != nil {
		return Point{}, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return Point{}, err
	}

	return resp.Point, nil
}

// GetBounds get bounding rectangle.
func (client *Tile38Client) GetBounds(key, objectID string) (Bounds, error) {
	var resp struct {
		Bounds Bounds `json:"bounds"`
	}

	b, err := client.Execute("GET", key, objectID, "BOUNDS")
	if err != nil {
		return Bounds{}, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return Bounds{}, err
	}

	return resp.Bounds, nil
}

// GetHash returns Geohash of object. Requires a precision of 1 to 22.
func (client *Tile38Client) GetHash(key, objectID string, precision int) (string, error) {
	var resp struct {
		Hash string `json:"hash"`
	}

	b, err := client.Execute("GET", key, objectID, "HASH", strconv.Itoa(precision))
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return "", err
	}

	return resp.Hash, nil
}
