package t38c

import (
	"strings"
)

func (client *Tile38Client) searchObjects(cmd, key, area string, opts []SearchOption) ([]*Object, error) {
	var resp struct {
		Fields  []string `json:"fields"`
		Objects []struct {
			ID     string  `json:"id"`
			Object *Object `json:"object"`
			Fields []int   `json:"fields"`
		} `json:"objects"`
	}

	command := buildSearchCommand(cmd, key, area, opts)
	if err := client.execute(command, &resp); err != nil {
		return nil, err
	}

	objects := make([]*Object, len(resp.Objects))
	haveFields := len(resp.Fields) > 0
	for idx, obj := range resp.Objects {
		o := obj.Object
		o.Tile38ID = obj.ID
		if haveFields {
			o.Fields = make(map[string]int)
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
func (client *Tile38Client) Intersects(key string, area SearchArea, opts ...SearchOption) ([]*Object, error) {
	return client.searchObjects("INTERSECTS", key, string(area), opts)
}

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func (client *Tile38Client) Within(key string, area SearchArea, opts ...SearchOption) ([]*Object, error) {
	return client.searchObjects("WITHIN", key, string(area), opts)
}

// Nearby searches a collection for objects that are close to a specified point.
func (client *Tile38Client) Nearby(key string, area NearbyArea, opts ...SearchOption) ([]*Object, error) {
	return client.searchObjects("NEARBY", key, string(area), opts)
}
