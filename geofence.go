package t38c

// GeofenceWithin return Within geofence request.
func (client *Client) GeofenceWithin(key string, area SearchArea) GeofenceQueryBuilder {
	return GeofenceQueryBuilder{
		client: client,
		cmd:    "WITHIN",
		key:    key,
		area:   Command(area),
	}
}

// GeofenceIntersects return Intersects geofence request.
func (client *Client) GeofenceIntersects(key string, area SearchArea) GeofenceQueryBuilder {
	return GeofenceQueryBuilder{
		client: client,
		cmd:    "INTERSECTS",
		key:    key,
		area:   Command(area),
	}
}

// GeofenceNearby return Nearby geofence request.
func (client *Client) GeofenceNearby(key string, lat, lon, meters float64) GeofenceQueryBuilder {
	return GeofenceQueryBuilder{
		client: client,
		cmd:    "NEARBY",
		key:    key,
		area:   NewCommand("POINT", floatString(lat), floatString(lon), floatString(meters)),
	}
}

// GeofenceRoam return roaming geofence request.
func (client *Client) GeofenceRoam(key, target, pattern string, meters int) GeofenceQueryBuilder {
	return GeofenceQueryBuilder{
		client:      client,
		isRoamQuery: true,
		key:         key,
		target:      target,
		pattern:     pattern,
		meters:      meters,
	}
}
