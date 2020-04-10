package geofence

import (
	"strconv"

	t38c "github.com/zerobounty/tile38-client"
)

var _ Requestable = (*RoamRequest)(nil)

// RoamRequest struct
type RoamRequest struct {
	Key           string
	Target        string
	Pattern       string
	Meters        int
	ObjectType    t38c.Command
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

	if len(req.ObjectType.Name) > 0 {
		args = append(args, req.ObjectType.Name)
		args = append(args, req.ObjectType.Args...)
	}

	args = append(args, []string{
		"ROAM", req.Target, req.Pattern, strconv.Itoa(req.Meters),
	}...)

	return t38c.NewCommand("NEARBY", args...)
}

// NewRoamRequest ...
func NewRoamRequest(key, target, pattern string, meters int) *RoamRequest {
	return &RoamRequest{
		Key:     key,
		Target:  target,
		Pattern: pattern,
		Meters:  meters,
	}
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
