package t38c

// GeofenceWithin return Within geofence request.
func (client *Client) GeofenceWithin(key string) GeofenceAreaSelector {
	return newGeofenceAreaSelector(client, "WITHIN", key)
}

// GeofenceIntersects return Intersects geofence request.
func (client *Client) GeofenceIntersects(key string) GeofenceAreaSelector {
	return newGeofenceAreaSelector(client, "INTERSECTS", key)
}

// GeofenceNearby return Nearby geofence request.
func (client *Client) GeofenceNearby(key string, lat, lon, meters float64) GeofenceQueryBuilder {
	area := newTileCmd("POINT", floatString(lat), floatString(lon), floatString(meters))
	return newGeofenceQueryBuilder(client, "NEARBY", key, area)
}

// GeofenceRoam return roaming geofence request.
func (client *Client) GeofenceRoam(key, target, pattern string, meters int) GeofenceQueryBuilder {
	return newGeofenceRoamQueryBuilder(client, key, target, pattern, meters)
}
