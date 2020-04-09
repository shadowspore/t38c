package t38c

import "strconv"

var _ GeofenceRequestable = (*GeofenceRoamRequest)(nil)

// GeofenceRoamRequest struct
type GeofenceRoamRequest struct {
	Key           string
	Target        string
	Pattern       string
	Meters        int
	ObjectType    Command
	DetectActions []GeofenceDetectAction
	Options       []SearchOption
}

// GeofenceCommand ...
func (req *GeofenceRoamRequest) GeofenceCommand() Command {
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

	return NewCommand("NEARBY", args...)
}

// NewFenceRoam ...
func NewFenceRoam(key, target, pattern string, meters int) *GeofenceRoamRequest {
	return &GeofenceRoamRequest{
		Key:     key,
		Target:  target,
		Pattern: pattern,
		Meters:  meters,
	}
}

// Actions ...
func (req *GeofenceRoamRequest) Actions(actions ...GeofenceDetectAction) *GeofenceRoamRequest {
	req.DetectActions = actions
	return req
}

// WithOptions ...
func (req *GeofenceRoamRequest) WithOptions(opts ...SearchOption) *GeofenceRoamRequest {
	req.Options = opts
	return req
}
