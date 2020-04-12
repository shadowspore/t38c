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

// Expire set a timeout on an id.
func (client *Tile38Client) Expire(key, objectID string, seconds int) error {
	_, err := client.Execute("EXPIRE", key, objectID, strconv.Itoa(seconds))
	return err
}

// FSet set the value for one or more fields of an id. Fields are double precision floating points.
// Normally, FSET will return an error if the field is being set on a non-existent id.
// However, the option XX can alter this behavior.
// Specifically, if called with XX option, FSET will return 0 when called on a non-existend id.
// Note that the non-existent key will still cause an error!
func (client *Tile38Client) FSet(key, objectID string, opts ...SetOption) error {
	var args []string = []string{
		key, objectID,
	}

	for _, opt := range opts {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}

	_, err := client.Execute("FSET", args...)
	return err
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

// JDel delete a value from a JSON document.
func (client *Tile38Client) JDel(key, objectID, path string) error {
	_, err := client.Execute("JDEL", key, objectID, path)
	return err
}

// JGet get a value from a JSON document.
func (client *Tile38Client) JGet(key, objectID, path string) (json.RawMessage, error) {
	return client.Execute("JGET", key, objectID, path)
}

// JSet set a value in a JSON document.
func (client *Tile38Client) JSet(key, objectID, path, value string) (json.RawMessage, error) {
	return client.Execute("JSET", key, objectID, path, value)
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

// PDel removes objects that match a specified pattern.
func (client *Tile38Client) PDel(key, pattern string) error {
	_, err := client.Execute("PDEL", key, pattern)
	return err
}

// Persist remove an existing timeout of an id.
func (client *Tile38Client) Persist(key, objectID string) error {
	_, err := client.Execute("PERSIST", key, objectID)
	return err
}

// Rename renames collection key to newkey.
// If newkey already exists, it will be deleted prior to renaming.
// Returns “OK” for success or “ERR” when key or newkey are actively being used by a geofence or webhook.
func (client *Tile38Client) Rename(key, newKey string) error {
	_, err := client.Execute("RENAME", key, newKey)
	return err
}

// RenameNX renames collection key to newkey, if it does not exist yet.
// If newkey already exists, this command does nothing.
// Returns 1 if key was renamed to newkey, 0 if newkey already existed,
// or “ERR” when key or newkey are actively being used by a geofence or webhook.
func (client *Tile38Client) RenameNX(key, newKey string) error {
	_, err := client.Execute("RENAMENX", key, newKey)
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

// Stats return stats for one or more keys.
func (client *Tile38Client) Stats(keys ...string) ([]KeyStats, error) {
	var resp struct {
		Stats []KeyStats `json:"stats"`
	}

	b, err := client.Execute("STATS", keys...)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return nil, err
	}

	return resp.Stats, nil
}

// TTL get a timeout on an id.
func (client *Tile38Client) TTL(key, objectID string) (int, error) {
	var resp struct {
		TTL int `json:"ttl"`
	}

	b, err := client.Execute("TTL", key, objectID)
	if err != nil {
		return -1, err
	}

	if err := json.Unmarshal(b, &resp); err != nil {
		return -1, err
	}

	return resp.TTL, nil
}
