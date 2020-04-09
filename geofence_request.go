package t38c

// GeofenceRequestable interface
// TODO: rename?
type GeofenceRequestable interface {
	GeofenceCommand() Command
}

var _ GeofenceRequestable = (*GeofenceRequest)(nil)

// GeofenceRequest struct
type GeofenceRequest struct {
	Cmd           string
	Key           string
	Area          SearchArea
	ObjectType    Command
	DetectActions []GeofenceDetectAction
	Options       []SearchOption
}

// GeofenceCommand ...
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

	if len(req.ObjectType.Name) > 0 {
		args = append(args, req.ObjectType.Name)
		args = append(args, req.ObjectType.Args...)
	}

	args = append(args, req.Area.Name)
	args = append(args, req.Area.Args...)

	return NewCommand(req.Cmd, args...)
}

// Actions ...
func (req *GeofenceRequest) Actions(actions ...GeofenceDetectAction) *GeofenceRequest {
	req.DetectActions = actions
	return req
}

// WithOptions ...
func (req *GeofenceRequest) WithOptions(opts ...SearchOption) *GeofenceRequest {
	req.Options = opts
	return req
}

// NewFenceReq ...
func NewFenceReq(cmd string, key string, area SearchArea) *GeofenceRequest {
	return &GeofenceRequest{
		Cmd:  cmd,
		Key:  key,
		Area: area,
	}
}
