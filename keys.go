package t38c

import (
	"context"
	"encoding/json"
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

// Keys struct
type Keys struct {
	client tile38Client
}

// Bounds returns the minimum bounding rectangle for all objects in a key.
func (ks *Keys) Bounds(ctx context.Context, key string) ([][][]float64, error) {
	var resp struct {
		Bounds geojson.Geometry `json:"bounds"`
	}

	err := ks.client.jExecute(ctx, &resp, "BOUNDS", key)
	if err != nil {
		return nil, err
	}

	return resp.Bounds.Polygon, nil
}

// Del remove a specified object.
func (ks *Keys) Del(ctx context.Context, key, objectID string) error {
	return ks.client.jExecute(ctx, nil, "DEL", key, objectID)
}

// Drop remove all objects from specified key.
func (ks *Keys) Drop(ctx context.Context, key string) error {
	return ks.client.jExecute(ctx, nil, "DROP", key)
}

// Expire set a timeout on an id.
func (ks *Keys) Expire(ctx context.Context, key, objectID string, seconds int) error {
	return ks.client.jExecute(ctx, nil, "EXPIRE", key, objectID, strconv.Itoa(seconds))
}

// FSet set the value for one or more fields of an id. Fields are double precision floating points.
// Normally, FSET will return an error if the field is being set on a non-existent id.
// However, the option XX can alter this behavior.
// Specifically, if called with XX option, FSET will return 0 when called on a non-existend id.
// Note that the non-existent key will still cause an error!
func (ks *Keys) FSet(key, objectID string) FSetQueryBuilder {
	return newFSetQueryBuilder(ks.client, key, objectID)
}

// Get returns object of an id.
func (ks *Keys) Get(key, objectID string) KeysGetQueryBuilder {
	return newKeysGetQueryBuilder(ks.client, key, objectID)
}

// JDel delete a value from a JSON document.
func (ks *Keys) JDel(ctx context.Context, key, objectID, path string) error {
	return ks.client.jExecute(ctx, nil, "JDEL", key, objectID, path)
}

// JGet get a value from a JSON document.
func (ks *Keys) JGet(ctx context.Context, key, objectID, path string) ([]byte, error) {
	var resp struct {
		Value json.RawMessage `json:"value"`
	}

	// cmd := newTileCmd("JGET", key, objectID, path)
	// if raw {
	// 	cmd.appendArgs("RAW")
	// }

	err := ks.client.jExecute(ctx, &resp, "JGET", key, objectID, path)
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}

// JSet set a value in a JSON document.
func (ks *Keys) JSet(key, objectID, path, value string) JSetQueryBuilder {
	return newJSetQueryBuilder(ks.client, key, objectID, path, value)
}

// Keys returns all keys matching pattern.
func (ks *Keys) Keys(ctx context.Context, pattern string) ([]string, error) {
	var resp struct {
		Keys []string `json:"keys"`
	}

	err := ks.client.jExecute(ctx, &resp, "KEYS", pattern)
	if err != nil {
		return nil, err
	}

	return resp.Keys, nil
}

// PDel removes objects that match a specified pattern.
func (ks *Keys) PDel(ctx context.Context, key, pattern string) error {
	return ks.client.jExecute(ctx, nil, "PDEL", key, pattern)
}

// Persist remove an existing timeout of an id.
func (ks *Keys) Persist(ctx context.Context, key, objectID string) error {
	return ks.client.jExecute(ctx, nil, "PERSIST", key, objectID)
}

// Rename renames collection key to newkey.
// If newkey already exists, it will be deleted prior to renaming.
// Returns “OK” for success or “ERR” when key or newkey are actively being used by a geofence or webhook.
func (ks *Keys) Rename(ctx context.Context, key, newKey string) error {
	return ks.client.jExecute(ctx, nil, "RENAME", key, newKey)
}

// RenameNX renames collection key to newkey, if it does not exist yet.
// If newkey already exists, this command does nothing.
// Returns 1 if key was renamed to newkey, 0 if newkey already existed,
// or “ERR” when key or newkey are actively being used by a geofence or webhook.
func (ks *Keys) RenameNX(ctx context.Context, key, newKey string) error {
	return ks.client.jExecute(ctx, nil, "RENAMENX", key, newKey)
}

// Set the value of an id. If a value is already associated to that key/id, it’ll be overwritten.
func (ks *Keys) Set(key, objectID string) SetAreaSelector {
	return newSetAreaSelector(ks.client, key, objectID)
}

// Stats return stats for one or more keys.
func (ks *Keys) Stats(ctx context.Context, keys ...string) ([]KeyStats, error) {
	var resp struct {
		Stats []KeyStats `json:"stats"`
	}

	err := ks.client.jExecute(ctx, &resp, "STATS", keys...)
	if err != nil {
		return nil, err
	}

	return resp.Stats, nil
}

// TTL get a timeout on an id.
func (ks *Keys) TTL(ctx context.Context, key, objectID string) (int, error) {
	var resp struct {
		TTL int `json:"ttl"`
	}

	err := ks.client.jExecute(ctx, &resp, "TTL", key, objectID)
	if err != nil {
		return -1, err
	}

	return resp.TTL, nil
}
