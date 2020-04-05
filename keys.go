package t38c

import (
	"encoding/json"
	"strconv"
	"strings"

	geojson "github.com/paulmach/go.geojson"
)

// Bounds returns the minimum bounding rectangle for all objects in a key.
func (client *Tile38Client) Bounds(key string) ([][][]float64, error) {
	var resp struct {
		Bounds geojson.Geometry `json:"bounds"`
	}
	if err := client.execute("BOUNDS "+key, &resp); err != nil {
		return nil, err
	}

	return resp.Bounds.Polygon, nil
}

// Del remove a specified object.
func (client *Tile38Client) Del(key, objectID string) error {
	err := client.execute("DEL "+key+" "+objectID, nil)
	return err
}

// Drop remove all objects from specified key.
func (client *Tile38Client) Drop(key string) error {
	err := client.execute("DROP "+key, nil)
	return err
}

// Set the value of an id. If a value is already associated to that key/id, it’ll be overwritten.
func (client *Tile38Client) Set(key, objectID string, area SetArea, opts ...SetOption) error {
	var sb strings.Builder
	sb.WriteString("SET " + key + " " + objectID)
	for _, opt := range opts {
		sb.WriteString(" " + string(opt))
	}
	sb.WriteString(" " + string(area))

	err := client.execute(sb.String(), nil)
	return err
}

// Keys returns all keys matching pattern.
func (client *Tile38Client) Keys(pattern string) ([]string, error) {
	var resp struct {
		Keys []string `json:"keys"`
	}
	err := client.execute("KEYS "+pattern, &resp)
	if err != nil {
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

	cmd := "GET " + key + " " + objectID
	if withFields {
		cmd += " WITHFIELDS"
	}

	err := client.execute(cmd, &resp)
	if err != nil {
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

	cmd := "GET " + key + " " + objectID + " POINT"
	err := client.execute(cmd, &resp)
	if err != nil {
		return Point{}, err
	}

	return resp.Point, nil
}

// GetBounds get bounding rectangle.
func (client *Tile38Client) GetBounds(key, objectID string) (Bounds, error) {
	var resp struct {
		Bounds Bounds `json:"bounds"`
	}

	cmd := "GET " + key + " " + objectID + " BOUNDS"
	err := client.execute(cmd, &resp)
	if err != nil {
		return Bounds{}, err
	}

	return resp.Bounds, nil
}

// GetHash returns Geohash of object. Requires a precision of 1 to 22.
func (client *Tile38Client) GetHash(key, objectID string, precision int) (string, error) {
	var resp struct {
		Hash string `json:"hash"`
	}

	cmd := "GET " + key + " HASH " + strconv.Itoa(precision) + " " + objectID
	err := client.execute(cmd, &resp)
	if err != nil {
		return "", err
	}

	return resp.Hash, nil
}
