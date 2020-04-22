package t38c

import (
	"context"
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
	Cmd           string
	Key           string
	Area          Command
	OutputFormat  OutputFormat
	DetectActions []DetectAction
	Options       []SearchOption
}

// GeofenceCommand build geofence command for tile38.
func (req *GeofenceRequest) GeofenceCommand() Command {
	var args []string
	args = append(args, req.Key)

	for _, opt := range req.Options {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}

	args = append(args, "FENCE")

	if len(req.DetectActions) > 0 {
		args = append(args, "DETECT")
		actions := ""
		first := true
		for _, action := range req.DetectActions {
			if !first {
				actions += ","
			}
			actions += string(action)
			first = false
		}
		args = append(args, actions)
	}

	if len(req.OutputFormat.Name) > 0 {
		args = append(args, req.OutputFormat.Name)
		args = append(args, req.OutputFormat.Args...)
	}

	args = append(args, req.Area.Name)
	args = append(args, req.Area.Args...)

	return NewCommand(req.Cmd, args...)
}

// Actions sets the geofence actions to receive.
// All actions used by default.
func (req *GeofenceRequest) Actions(actions ...DetectAction) *GeofenceRequest {
	req.DetectActions = actions
	return req
}

// WithOptions sets the optional parameters for request.
func (req *GeofenceRequest) WithOptions(opts ...SearchOption) *GeofenceRequest {
	req.Options = opts
	return req
}

// Format set geofence GeofenceResponse format.
func (req *GeofenceRequest) Format(fmt OutputFormat) *GeofenceRequest {
	req.OutputFormat = fmt
	return req
}

// GeofenceWithin return Within geofence request.
func GeofenceWithin(key string, area SearchArea) *GeofenceRequest {
	return &GeofenceRequest{
		Cmd:  "WITHIN",
		Key:  key,
		Area: Command(area),
	}
}

// GeofenceIntersects return Intersects geofence request.
func GeofenceIntersects(key string, area SearchArea) *GeofenceRequest {
	return &GeofenceRequest{
		Cmd:  "INTERSECTS",
		Key:  key,
		Area: Command(area),
	}
}

// GeofenceNearby return Nearby geofence request.
func GeofenceNearby(key string, lat, lon, meters float64) *GeofenceRequest {
	return &GeofenceRequest{
		Cmd:  "NEARBY",
		Key:  key,
		Area: NewCommand("POINT", floatString(lat), floatString(lon), floatString(meters)),
	}
}

var _ GeofenceRequestable = (*RoamGeofenceRequest)(nil)

// RoamGeofenceRequest represents a roaming geofence request.
type RoamGeofenceRequest struct {
	Key           string
	Target        string
	Pattern       string
	Meters        int
	OutputFormat  OutputFormat
	DetectActions []DetectAction
	Options       []SearchOption
}

// GeofenceCommand build geofence command for tile38.
func (req *RoamGeofenceRequest) GeofenceCommand() Command {
	var args []string
	args = append(args, req.Key)

	for _, opt := range req.Options {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}

	args = append(args, "FENCE")

	if len(req.DetectActions) > 0 {
		args = append(args, "DETECT")
		actions := ""
		first := true
		for _, action := range req.DetectActions {
			if !first {
				actions += ","
			}
			actions += string(action)
			first = false
		}
		args = append(args, actions)
	}

	if len(req.OutputFormat.Name) > 0 {
		args = append(args, req.OutputFormat.Name)
		args = append(args, req.OutputFormat.Args...)
	}

	args = append(args, []string{
		"ROAM", req.Target, req.Pattern, strconv.Itoa(req.Meters),
	}...)

	return NewCommand("NEARBY", args...)
}

// Actions sets the geofence actions to receive.
// All actions used by default.
func (req *RoamGeofenceRequest) Actions(actions ...DetectAction) *RoamGeofenceRequest {
	req.DetectActions = actions
	return req
}

// WithOptions sets the optional parameters for request.
func (req *RoamGeofenceRequest) WithOptions(opts ...SearchOption) *RoamGeofenceRequest {
	req.Options = opts
	return req
}

// Format set geofence GeofenceResponse format.
func (req *RoamGeofenceRequest) Format(fmt OutputFormat) *RoamGeofenceRequest {
	req.OutputFormat = fmt
	return req
}

// GeofenceRoam return roaming geofence request.
func GeofenceRoam(key, target, pattern string, meters int) *RoamGeofenceRequest {
	return &RoamGeofenceRequest{
		Key:     key,
		Target:  target,
		Pattern: pattern,
		Meters:  meters,
	}
}

// Fence execute geofence command.
func (client *Client) Fence(ctx context.Context, req GeofenceRequestable) (chan GeofenceResponse, error) {
	cmd := req.GeofenceCommand()
	events, err := client.ExecuteStream(ctx, cmd.Name, cmd.Args...)
	if err != nil {
		return nil, err
	}

	return unmarshalEvents(events)
}
