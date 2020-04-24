package t38c

import "strings"

// GeofenceParams optional params
type GeofenceParams struct {
	OutputFormat   OutputFormat
	DetectActions  []DetectAction
	NotifyCommands []NotifyCommand
	Options        []SearchOption
}

func (params *GeofenceParams) args() []string {
	args := []string{}
	for _, opt := range params.Options {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}

	args = append(args, "FENCE")
	if len(params.DetectActions) > 0 {
		args = append(args, "DETECT")
		actions := make([]string, len(params.DetectActions))
		for i := range params.DetectActions {
			actions[i] = string(params.DetectActions[i])
		}
		args = append(args, strings.Join(actions, ","))
	}

	if len(params.NotifyCommands) > 0 {
		args = append(args, "COMMANDS")
		actions := make([]string, len(params.NotifyCommands))
		for i := range params.NotifyCommands {
			actions[i] = string(params.NotifyCommands[i])
		}
		args = append(args, strings.Join(actions, ","))
	}

	if len(params.OutputFormat.Name) > 0 {
		args = append(args, params.OutputFormat.Name)
		args = append(args, params.OutputFormat.Args...)
	}

	return args
}

// GeofenceOption ...
type GeofenceOption func(*GeofenceParams)

// Actions sets the geofence actions.
// All actions used by default.
func Actions(actions ...DetectAction) GeofenceOption {
	return func(params *GeofenceParams) {
		params.DetectActions = actions
	}
}

// Commands sets the geofence commands.
func Commands(notifyCommands ...NotifyCommand) GeofenceOption {
	return func(params *GeofenceParams) {
		params.NotifyCommands = notifyCommands
	}
}

// SearchOptions sets the optional parameters for request.
func SearchOptions(opts ...SearchOption) GeofenceOption {
	return func(params *GeofenceParams) {
		params.Options = opts
	}
}

// Format set geofence GeofenceResponse format.
func Format(fmt OutputFormat) GeofenceOption {
	return func(params *GeofenceParams) {
		params.OutputFormat = fmt
	}
}

func getParams(opts ...GeofenceOption) *GeofenceParams {
	if len(opts) == 0 {
		return nil
	}

	params := &GeofenceParams{}
	for _, opt := range opts {
		opt(params)
	}

	return params
}
