package geofence

import (
	"strconv"

	t38c "github.com/lostpeer/tile38-client"
)

var _ Requestable = (*RoamRequest)(nil)

// RoamRequest struct
type RoamRequest struct {
	Key           string
	Target        string
	Pattern       string
	Meters        int
	OutputFormat  t38c.OutputFormat
	DetectActions []DetectAction
	Options       []t38c.SearchOption
}

// GeofenceCommand ...
func (req *RoamRequest) GeofenceCommand() t38c.Command {
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

	return t38c.NewCommand("NEARBY", args...)
}

// Actions ...
func (req *RoamRequest) Actions(actions ...DetectAction) *RoamRequest {
	req.DetectActions = actions
	return req
}

// WithOptions ...
func (req *RoamRequest) WithOptions(opts ...t38c.SearchOption) *RoamRequest {
	req.Options = opts
	return req
}

// Format ...
func (req *RoamRequest) Format(fmt t38c.OutputFormat) *RoamRequest {
	req.OutputFormat = fmt
	return req
}

// Roam ...
func Roam(key, target, pattern string, meters int) *RoamRequest {
	return &RoamRequest{
		Key:     key,
		Target:  target,
		Pattern: pattern,
		Meters:  meters,
	}
}
