package t38c

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func (client *Client) Within(key string) InwAreaSelector {
	return newInwAreaSelector(client, "WITHIN", key)
}

// Intersects searches a collection for objects that intersect a specified bounding area.
func (client *Client) Intersects(key string) InwAreaSelector {
	return newInwAreaSelector(client, "INTERSECTS", key)
}

// Nearby command searches a collection for objects that are close to a specified point.
// The KNN algorithm is used instead of the standard overlap+Haversine algorithm,
// sorting the results in order of ascending distance from that point, i.e., nearest first.
func (client *Client) Nearby(key string, lat, lon, meters float64) InwQueryBuilder {
	area := newTileCmd("POINT", floatString(lat), floatString(lon), floatString(meters))
	return newInwQueryBuilder(client, "NEARBY", key, area)
}

// Search iterates though a key’s string values.
func (client *Client) Search(key string) SearchQueryBuilder {
	return SearchQueryBuilder{
		client: client,
		key:    key,
	}
}

// Scan incrementally iterates though a key.
func (client *Client) Scan(key string) ScanQueryBuilder {
	return ScanQueryBuilder{
		client: client,
		key:    key,
	}
}
