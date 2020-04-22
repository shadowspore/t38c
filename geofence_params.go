package t38c

import "strings"

// optional params
type geofenceParams struct {
	OutputFormat   OutputFormat
	DetectActions  []DetectAction
	NotifyCommands []NotifyCommand
	Options        []SearchOption
}

func (params *geofenceParams) args() []string {
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

// Actions sets the geofence actions.
// All actions used by default.
func Actions(actions ...DetectAction) func(*geofenceParams) {
	return func(params *geofenceParams) {
		params.DetectActions = actions
	}
}

// Commands sets the geofence commands.
func Commands(notifyCommands ...NotifyCommand) func(*geofenceParams) {
	return func(params *geofenceParams) {
		params.NotifyCommands = notifyCommands
	}
}

// GeofenceOptions sets the optional parameters for request.
func GeofenceOptions(opts ...SearchOption) func(*geofenceParams) {
	return func(params *geofenceParams) {
		params.Options = opts
	}
}

// Format set geofence GeofenceResponse format.
func Format(fmt OutputFormat) func(*geofenceParams) {
	return func(params *geofenceParams) {
		params.OutputFormat = fmt
	}
}

func getParams(opts ...func(*geofenceParams)) *geofenceParams {
	if len(opts) == 0 {
		return nil
	}

	params := &geofenceParams{}
	for _, opt := range opts {
		opt(params)
	}

	return params
}
