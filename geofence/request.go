package geofence

import t38c "github.com/lostpeer/tile38-client"

// Requestable interface
// TODO: rename?
type Requestable interface {
	GeofenceCommand() t38c.Command
}

var _ Requestable = (*Request)(nil)

// Request struct
type Request struct {
	Cmd           string
	Key           string
	Area          t38c.SearchArea
	ObjectType    t38c.Command
	DetectActions []DetectAction
	Options       []t38c.SearchOption
}

// Command ...
func (req *Request) GeofenceCommand() t38c.Command {
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

	return t38c.NewCommand(req.Cmd, args...)
}

// Actions ...
func (req *Request) Actions(actions ...DetectAction) *Request {
	req.DetectActions = actions
	return req
}

// WithOptions ...
func (req *Request) WithOptions(opts ...t38c.SearchOption) *Request {
	req.Options = opts
	return req
}

// NewFenceReq ...
func NewFenceReq(cmd string, key string, area t38c.SearchArea) *Request {
	return &Request{
		Cmd:  cmd,
		Key:  key,
		Area: area,
	}
}
