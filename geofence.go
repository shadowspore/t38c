package t38c

import (
	"context"
	"fmt"
	"log"
)

// GeofenceWithin return Within geofence request.
func GeofenceWithin(key string, area SearchArea) geofenceQueryBuilder {
	return geofenceQueryBuilder{
		cmd:       "WITHIN",
		key:       key,
		area:      Command(area),
	}
}

// GeofenceIntersects return Intersects geofence request.
func GeofenceIntersects(key string, area SearchArea) geofenceQueryBuilder {
	return geofenceQueryBuilder{
		cmd:       "INTERSECTS",
		key:       key,
		area:      Command(area),
	}
}

// GeofenceNearby return Nearby geofence request.
func GeofenceNearby(key string, lat, lon, meters float64) geofenceQueryBuilder {
	return geofenceQueryBuilder{
		cmd:       "NEARBY",
		key:       key,
		area:      NewCommand("POINT", floatString(lat), floatString(lon), floatString(meters)),
	}
}

// GeofenceRoam return roaming geofence request.
func GeofenceRoam(key, target, pattern string, meters int) geofenceQueryBuilder {
	return geofenceQueryBuilder{
		isRoamQuery: true,
		key:         key,
		target:      target,
		pattern:     pattern,
		meters:      meters,
	}
}

// Fence execute geofence command.
func (client *Client) Fence(ctx context.Context, query geofenceQueryBuilder) (chan GeofenceResponse, error) {
	cmd := query.Cmd()
	events, err := client.ExecuteStream(ctx, cmd.Name, cmd.Args...)
	if err != nil {
		return nil, fmt.Errorf("command: %s: %v", cmd, err)
	}

	if client.debug {
		log.Printf("[%s]: ok\n", cmd)
	}

	return unmarshalEvents(events)
}
