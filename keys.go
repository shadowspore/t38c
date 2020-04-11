package t38c

import (
	"encoding/json"

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

// Get returns object of an id.
func (client *Tile38Client) Get(key, objectID string, withFields bool) (*GetResponse, error) {
	var resp *GetResponse

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

	err = json.Unmarshal(b, &resp)
	return resp, err
}
