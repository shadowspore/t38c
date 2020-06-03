package t38c

// Geofence struct
type Geofence struct {
	client tile38Client
}

// Within geofence
func (gf *Geofence) Within(key string) GeofenceAreaSelector {
	return newGeofenceAreaSelector(gf.client, "WITHIN", key)
}

// Intersects geofence
func (gf *Geofence) Intersects(key string) GeofenceAreaSelector {
	return newGeofenceAreaSelector(gf.client, "INTERSECTS", key)
}

// Nearby geofence
func (gf *Geofence) Nearby(key string, lat, lon, meters float64) GeofenceQueryBuilder {
	area := newTileCmd("POINT", floatString(lat), floatString(lon), floatString(meters))
	return newGeofenceQueryBuilder(gf.client, "NEARBY", key, area)
}

// Roam geofence
func (gf *Geofence) Roam(key, target, pattern string, meters int) GeofenceQueryBuilder {
	return newGeofenceRoamQueryBuilder(gf.client, key, target, pattern, meters)
}
