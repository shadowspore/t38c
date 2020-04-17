package t38c

import (
	"encoding/json"
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

// Bounds returns the minimum bounding rectangle for all objects in a key.
func (client *Client) Bounds(key string) ([][][]float64, error) {
	var resp struct {
		Bounds geojson.Geometry `json:"bounds"`
	}

	err := client.jExecute(&resp, "BOUNDS", key)
	if err != nil {
		return nil, err
	}

	return resp.Bounds.Polygon, nil
}

// Del remove a specified object.
func (client *Client) Del(key, objectID string) error {
	return client.jExecute(nil, "DEL", key)
}

// Drop remove all objects from specified key.
func (client *Client) Drop(key string) error {
	return client.jExecute(nil, "DROP", key)
}

// Expire set a timeout on an id.
func (client *Client) Expire(key, objectID string, seconds int) error {
	return client.jExecute(nil, "EXPIRE", key, objectID, strconv.Itoa(seconds))
}

// FSet set the value for one or more fields of an id. Fields are double precision floating points.
// Normally, FSET will return an error if the field is being set on a non-existent id.
// However, the option XX can alter this behavior.
// Specifically, if called with XX option, FSET will return 0 when called on a non-existend id.
// Note that the non-existent key will still cause an error!
func (client *Client) FSet(key, objectID string, opts ...SetOption) error {
	var args []string = []string{
		key, objectID,
	}

	for _, opt := range opts {
		if opt.Name == "FIELD" {
			args = append(args, opt.Args...)
			continue
		}

		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}

	return client.jExecute(nil, "FSET", args...)
}

// Get returns object of an id.
func (client *Client) Get(key, objectID string, withFields bool) (resp *GetResponse, err error) {
	var args []string
	args = append(args, key)
	args = append(args, objectID)
	if withFields {
		args = append(args, "WITHFIELDS")
	}

	err = client.jExecute(resp, "GET", args...)
	return
}

// JDel delete a value from a JSON document.
func (client *Client) JDel(key, objectID, path string) error {
	return client.jExecute(nil, "JDEL", key, objectID, path)
}

// JGet get a value from a JSON document.
func (client *Client) JGet(key, objectID, path string) (json.RawMessage, error) {
	return client.Execute("JGET", key, objectID, path)
}

// JSet set a value in a JSON document.
func (client *Client) JSet(key, objectID, path, value string) (json.RawMessage, error) {
	return client.Execute("JSET", key, objectID, path, value)
}

// Keys returns all keys matching pattern.
func (client *Client) Keys(pattern string) ([]string, error) {
	var resp struct {
		Keys []string `json:"keys"`
	}

	err := client.jExecute(&resp, "KEYS", pattern)
	if err != nil {
		return nil, err
	}

	return resp.Keys, nil
}

// PDel removes objects that match a specified pattern.
func (client *Client) PDel(key, pattern string) error {
	return client.jExecute(nil, "PDEL", key, pattern)
}

// Persist remove an existing timeout of an id.
func (client *Client) Persist(key, objectID string) error {
	return client.jExecute(nil, "PERSIST", key, objectID)
}

// Rename renames collection key to newkey.
// If newkey already exists, it will be deleted prior to renaming.
// Returns “OK” for success or “ERR” when key or newkey are actively being used by a geofence or webhook.
func (client *Client) Rename(key, newKey string) error {
	return client.jExecute(nil, "RENAME", key, newKey)
}

// RenameNX renames collection key to newkey, if it does not exist yet.
// If newkey already exists, this command does nothing.
// Returns 1 if key was renamed to newkey, 0 if newkey already existed,
// or “ERR” when key or newkey are actively being used by a geofence or webhook.
func (client *Client) RenameNX(key, newKey string) error {
	return client.jExecute(nil, "RENAMENX", key, newKey)
}

// Set the value of an id. If a value is already associated to that key/id, it’ll be overwritten.
func (client *Client) Set(key, objectID string, area SetArea, opts ...SetOption) error {
	var args []string
	args = append(args, key)
	args = append(args, objectID)
	for _, opt := range opts {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}
	args = append(args, area.Name)
	args = append(args, area.Args...)

	return client.jExecute(nil, "SET", args...)
}

// Stats return stats for one or more keys.
func (client *Client) Stats(keys ...string) ([]KeyStats, error) {
	var resp struct {
		Stats []KeyStats `json:"stats"`
	}

	err := client.jExecute(&resp, "STATS", keys...)
	if err != nil {
		return nil, err
	}

	return resp.Stats, nil
}

// TTL get a timeout on an id.
func (client *Client) TTL(key, objectID string) (int, error) {
	var resp struct {
		TTL int `json:"ttl"`
	}

	err := client.jExecute(&resp, "TTL", key, objectID)
	if err != nil {
		return -1, err
	}

	return resp.TTL, nil
}
