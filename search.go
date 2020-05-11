package t38c

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func (client *Client) Within(key string, area SearchArea) SearchQueryBuilder {
	return SearchQueryBuilder{
		client: client,
		cmd:    "WITHIN",
		key:    key,
		area:   Command(area),
	}
}

// Intersects searches a collection for objects that intersect a specified bounding area.
func (client *Client) Intersects(key string, area SearchArea) SearchQueryBuilder {
	return SearchQueryBuilder{
		client: client,
		cmd:    "INTERSECTS",
		key:    key,
		area:   Command(area),
	}
}

// Nearby command searches a collection for objects that are close to a specified point.
// The KNN algorithm is used instead of the standard overlap+Haversine algorithm,
// sorting the results in order of ascending distance from that point, i.e., nearest first.
func (client *Client) Nearby(key string, lat, lon, meters float64) SearchQueryBuilder {
	return SearchQueryBuilder{
		client: client,
		cmd:    "NEARBY",
		key:    key,
		area:   NewCommand("POINT", floatString(lat), floatString(lon), floatString(meters)),
	}
}

// Search iterates though a key’s string values.
func (client *Client) Search(key string) SearchQueryBuilder {
	return SearchQueryBuilder{
		client: client,
		cmd:    "SEARCH",
		key:    key,
	}
}

// Scan incrementally iterates though a key.
func (client *Client) Scan(key string) SearchQueryBuilder {
	return SearchQueryBuilder{
		client: client,
		cmd:    "SCAN",
		key:    key,
	}
}
