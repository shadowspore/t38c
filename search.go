package t38c

import (
	"strings"
)

type searchResponse struct {
	Fields  []string `json:"fields"`
	Objects []struct {
		ID     string  `json:"id"`
		Object *Object `json:"object"`
		Fields []int   `json:"fields"`
	} `json:"objects"`
}

func (sr searchResponse) getObjects() []*Object {
	objects := make([]*Object, len(sr.Objects))
	haveFields := len(sr.Fields) > 0
	for idx, obj := range sr.Objects {
		o := obj.Object
		o.Tile38ID = obj.ID
		if haveFields {
			o.Fields = make(map[string]int)
			for fieldIdx, field := range sr.Fields {
				o.Fields[field] = obj.Fields[fieldIdx]
			}
		}

		objects[idx] = o
	}

	return objects
}

// Intersects searches a collection for objects that intersect a specified bounding area.
func (client *Tile38Client) Intersects(key string, area SearchArea, opts ...SearchOption) ([]*Object, error) {
	var resp searchResponse
	cmd := buildSearchCommand("INTERSECTS", key, string(area), opts)
	err := client.execute(cmd, &resp)
	if err != nil {
		return nil, err
	}

	return resp.getObjects(), nil
}

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func (client *Tile38Client) Within(key string, area SearchArea, opts ...SearchOption) ([]*Object, error) {
	var resp searchResponse
	cmd := buildSearchCommand("WITHIN", key, string(area), opts)
	err := client.execute(cmd, &resp)
	if err != nil {
		return nil, err
	}

	return resp.getObjects(), nil
}

// Nearby searches a collection for objects that are close to a specified point.
func (client *Tile38Client) Nearby(key string, area NearbyArea, opts ...SearchOption) ([]*Object, error) {
	var resp searchResponse
	cmd := buildSearchCommand("NEARBY", key, string(area), opts)
	err := client.execute(cmd, &resp)
	if err != nil {
		return nil, err
	}

	return resp.getObjects(), nil
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
