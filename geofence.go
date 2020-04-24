package t38c

import (
	"context"
	"fmt"
	"log"
	"strconv"
)

// GeofenceRequestable interface.
// A geofence is a virtual boundary that can detect when an object enters or exits the area.
// This boundary can be a radius or any search area format, such as a bounding box, GeoJSON object, etc.
type GeofenceRequestable interface {
	GeofenceCommand() Command
}

var _ GeofenceRequestable = (*GeofenceRequest)(nil)

// GeofenceRequest represents a geofence request.
type GeofenceRequest struct {
	*GeofenceParams
	Cmd  string
	Key  string
	Area Command
}

// GeofenceCommand build geofence command for tile38.
func (req *GeofenceRequest) GeofenceCommand() Command {
	var args []string
	args = append(args, req.Key)
	if req.GeofenceParams != nil {
		args = append(args, req.GeofenceParams.args()...)
	}

	args = append(args, req.Area.Name)
	args = append(args, req.Area.Args...)

	return NewCommand(req.Cmd, args...)
}

// GeofenceWithin return Within geofence request.
func GeofenceWithin(key string, area SearchArea, opts ...GeofenceOption) *GeofenceRequest {
	return &GeofenceRequest{
		Cmd:            "WITHIN",
		Key:            key,
		Area:           Command(area),
		GeofenceParams: getParams(opts...),
	}
}

// GeofenceIntersects return Intersects geofence request.
func GeofenceIntersects(key string, area SearchArea, opts ...GeofenceOption) *GeofenceRequest {
	return &GeofenceRequest{
		Cmd:            "INTERSECTS",
		Key:            key,
		Area:           Command(area),
		GeofenceParams: getParams(opts...),
	}
}

// GeofenceNearby return Nearby geofence request.
func GeofenceNearby(key string, lat, lon, meters float64, opts ...GeofenceOption) *GeofenceRequest {
	return &GeofenceRequest{
		Cmd:            "NEARBY",
		Key:            key,
		Area:           NewCommand("POINT", floatString(lat), floatString(lon), floatString(meters)),
		GeofenceParams: getParams(opts...),
	}
}

var _ GeofenceRequestable = (*RoamGeofenceRequest)(nil)

// RoamGeofenceRequest represents a roaming geofence request.
type RoamGeofenceRequest struct {
	*GeofenceParams
	Key     string
	Target  string
	Pattern string
	Meters  int
}

// GeofenceCommand build geofence command for tile38.
func (req *RoamGeofenceRequest) GeofenceCommand() Command {
	var args []string
	args = append(args, req.Key)
	if req.GeofenceParams != nil {
		args = append(args, req.GeofenceParams.args()...)
	}

	args = append(args, []string{
		"ROAM", req.Target, req.Pattern, strconv.Itoa(req.Meters),
	}...)

	return NewCommand("NEARBY", args...)
}

// GeofenceRoam return roaming geofence request.
func GeofenceRoam(key, target, pattern string, meters int, opts ...GeofenceOption) *RoamGeofenceRequest {
	return &RoamGeofenceRequest{
		Key:            key,
		Target:         target,
		Pattern:        pattern,
		Meters:         meters,
		GeofenceParams: getParams(opts...),
	}
}

// Fence execute geofence command.
func (client *Client) Fence(ctx context.Context, req GeofenceRequestable) (chan GeofenceResponse, error) {
	cmd := req.GeofenceCommand()
	events, err := client.ExecuteStream(ctx, cmd.Name, cmd.Args...)
	if err != nil {
		return nil, fmt.Errorf("command: %s: %v", cmd, err)
	}

	if client.debug {
		log.Printf("[%s]: ok\n", cmd)
	}

	return unmarshalEvents(events)
}
