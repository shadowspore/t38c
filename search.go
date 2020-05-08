package t38c

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func Within(key string, area SearchArea) searchQueryBuilder {
	return searchQueryBuilder{
		cmd:  "WITHIN",
		key:  key,
		area: Command(area),
	}
}

// Intersects searches a collection for objects that intersect a specified bounding area.
func Intersects(key string, area SearchArea) searchQueryBuilder {
	return searchQueryBuilder{
		cmd:  "INTERSECTS",
		key:  key,
		area: Command(area),
	}
}

// Nearby command searches a collection for objects that are close to a specified point.
// The KNN algorithm is used instead of the standard overlap+Haversine algorithm,
// sorting the results in order of ascending distance from that point, i.e., nearest first.
func Nearby(key string, lat, lon, meters float64) searchQueryBuilder {
	return searchQueryBuilder{
		cmd:  "NEARBY",
		key:  key,
		area: NewCommand("POINT", floatString(lat), floatString(lon), floatString(meters)),
	}
}

// Search iterates though a key’s string values.
func Search(key string) searchQueryBuilder {
	return searchQueryBuilder{
		cmd: "SEARCH",
		key: key,
	}
}

// Scan incrementally iterates though a key.
func Scan(key string) searchQueryBuilder {
	return searchQueryBuilder{
		cmd: "SCAN",
		key: key,
	}
}

// Search execute a search request.
func (client *Client) Search(query searchQueryBuilder) (*SearchResponse, error) {
	cmd := query.Cmd()
	resp := &SearchResponse{}
	err := client.jExecute(&resp, cmd.Name, cmd.Args...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
