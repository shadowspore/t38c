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
	Area          t38c.Command
	OutputFormat  t38c.OutputFormat
	DetectActions []DetectAction
	Options       []t38c.SearchOption
}

// GeofenceCommand ...
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

	if len(req.OutputFormat.Name) > 0 {
		args = append(args, req.OutputFormat.Name)
		args = append(args, req.OutputFormat.Args...)
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

// Format ...
func (req *Request) Format(fmt t38c.OutputFormat) *Request {
	req.OutputFormat = fmt
	return req
}

// Within ...
func Within(key string, area t38c.SearchArea) *Request {
	return &Request{
		Cmd:  "WITHIN",
		Key:  key,
		Area: t38c.Command(area),
	}
}

// Intersects ...
func Intersects(key string, area t38c.SearchArea) *Request {
	return &Request{
		Cmd:  "INTERSECTS",
		Key:  key,
		Area: t38c.Command(area),
	}
}

// Nearby ...
func Nearby(key string, area t38c.NearbyArea) *Request {
	return &Request{
		Cmd:  "NEARBY",
		Key:  key,
		Area: t38c.Command(area),
	}
}
