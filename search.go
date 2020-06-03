package t38c

// Search struct
type Search struct {
	client tile38Client
}

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func (search *Search) Within(key string) InwAreaSelector {
	return newInwAreaSelector(search.client, "WITHIN", key)
}

// Intersects searches a collection for objects that intersect a specified bounding area.
func (search *Search) Intersects(key string) InwAreaSelector {
	return newInwAreaSelector(search.client, "INTERSECTS", key)
}

// Nearby command searches a collection for objects that are close to a specified point.
// The KNN algorithm is used instead of the standard overlap+Haversine algorithm,
// sorting the results in order of ascending distance from that point, i.e., nearest first.
func (search *Search) Nearby(key string, lat, lon, meters float64) InwQueryBuilder {
	area := newTileCmd("POINT", floatString(lat), floatString(lon), floatString(meters))
	return newInwQueryBuilder(search.client, "NEARBY", key, area)
}

// Search iterates though a key’s string values.
func (search *Search) Search(key string) SearchQueryBuilder {
	return newSearchQueryBuilder(search.client, key)
}

// Scan incrementally iterates though a key.
func (search *Search) Scan(key string) ScanQueryBuilder {
	return newScanQueryBuilder(search.client, key)
}
