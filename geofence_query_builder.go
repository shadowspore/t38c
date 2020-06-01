package t38c

import (
	"context"
	"strconv"
	"strings"
)

// GeofenceQueryBuilder optional params
type GeofenceQueryBuilder struct {
	client         *Client
	isRoamQuery    bool
	cmd            string
	key            string
	area           tileCmd
	target         string
	pattern        string
	meters         int
	outputFormat   OutputFormat
	detectActions  []DetectAction
	notifyCommands []NotifyCommand
	searchOpts     []tileCmd
}

func newGeofenceQueryBuilder(client *Client, cmd, key string, area tileCmd) GeofenceQueryBuilder {
	return GeofenceQueryBuilder{
		client: client,
		cmd:    cmd,
		key:    key,
		area:   area,
	}
}

func newGeofenceRoamQueryBuilder(client *Client, key, target, pattern string, meters int) GeofenceQueryBuilder {
	return GeofenceQueryBuilder{
		client:      client,
		cmd:         "NEARBY",
		isRoamQuery: true,
		key:         key,
		target:      target,
		pattern:     pattern,
		meters:      meters,
	}
}

func (query GeofenceQueryBuilder) toCmd() tileCmd {
	cmd := newTileCmd(query.cmd, query.key)
	for _, opt := range query.searchOpts {
		cmd.appendArgs(opt.Name, opt.Args...)
	}

	cmd.appendArgs("FENCE")
	if len(query.detectActions) > 0 {
		actions := make([]string, len(query.detectActions))
		for i := range query.detectActions {
			actions[i] = string(query.detectActions[i])
		}
		cmd.appendArgs("DETECT", strings.Join(actions, ","))
	}

	if len(query.notifyCommands) > 0 {
		actions := make([]string, len(query.notifyCommands))
		for i := range query.notifyCommands {
			actions[i] = string(query.notifyCommands[i])
		}
		cmd.appendArgs("COMMANDS", strings.Join(actions, ","))
	}

	if len(query.outputFormat.Name) > 0 {
		cmd.appendArgs(query.outputFormat.Name, query.outputFormat.Args...)
	}

	if query.isRoamQuery {
		cmd.appendArgs("ROAM", query.target, query.pattern, strconv.Itoa(query.meters))
		return cmd
	}

	cmd.appendArgs(query.area.Name, query.area.Args...)
	return cmd
}

// Do cmd
func (query GeofenceQueryBuilder) Do(ctx context.Context, handler func(*GeofenceEvent)) error {
	cmd := query.toCmd()
	return query.client.ExecuteStream(ctx, rawEventHandler(handler), cmd.Name, cmd.Args...)
}

// Actions sets the geofence actions.
// All actions used by default.
func (query GeofenceQueryBuilder) Actions(actions ...DetectAction) GeofenceQueryBuilder {
	query.detectActions = append(query.detectActions, actions...)
	return query
}

// Commands sets the geofence commands.
func (query GeofenceQueryBuilder) Commands(notifyCommands ...NotifyCommand) GeofenceQueryBuilder {
	query.notifyCommands = append(query.notifyCommands, notifyCommands...)
	return query
}

// Asc order. Only for SEARCH and SCAN commands.
func (query GeofenceQueryBuilder) Asc() GeofenceQueryBuilder {
	query.searchOpts = append(query.searchOpts, newTileCmd("ASC"))
	return query
}

// Desc order. Only for SEARCH and SCAN commands.
func (query GeofenceQueryBuilder) Desc() GeofenceQueryBuilder {
	query.searchOpts = append(query.searchOpts, newTileCmd("DESC"))
	return query
}

// NoFields tells the server that you do not want field values returned with the search results.
func (query GeofenceQueryBuilder) NoFields() GeofenceQueryBuilder {
	query.searchOpts = append(query.searchOpts, newTileCmd("NOFIELDS"))
	return query
}

// Clip tells the server to clip intersecting objects by the bounding box area of the search.
// It can only be used with these area formats: BOUNDS, TILE, QUADKEY, HASH.
func (query GeofenceQueryBuilder) Clip() GeofenceQueryBuilder {
	query.searchOpts = append(query.searchOpts, newTileCmd("CLIP"))
	return query
}

// Distance allows to return between objects. Only for NEARBY tileCmd.
func (query GeofenceQueryBuilder) Distance() GeofenceQueryBuilder {
	query.searchOpts = append(query.searchOpts, newTileCmd("DISTANCE"))
	return query
}

// Cursor is used to iterate though many objects from the search results.
// An iteration begins when the CURSOR is set to Zero or not included with the request,
// and completes when the cursor returned by the server is Zero.
func (query GeofenceQueryBuilder) Cursor(cursor int) GeofenceQueryBuilder {
	query.searchOpts = append(query.searchOpts, newTileCmd("CURSOR", strconv.Itoa(cursor)))
	return query
}

// Limit can be used to limit the number of objects returned for a single search request.
func (query GeofenceQueryBuilder) Limit(limit int) GeofenceQueryBuilder {
	query.searchOpts = append(query.searchOpts, newTileCmd("LIMIT", strconv.Itoa(limit)))
	return query
}

// Sparse will distribute the results of a search evenly across the requested area.
func (query GeofenceQueryBuilder) Sparse(sparse int) GeofenceQueryBuilder {
	query.searchOpts = append(query.searchOpts, newTileCmd("SPARSE", strconv.Itoa(sparse)))
	return query
}

// Where allows for filtering out results based on field values.
func (query GeofenceQueryBuilder) Where(field string, min, max float64) GeofenceQueryBuilder {
	query.searchOpts = append(query.searchOpts, newTileCmd("WHERE", field, floatString(min), floatString(max)))
	return query
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func (query GeofenceQueryBuilder) Wherein(field string, values ...float64) GeofenceQueryBuilder {
	cmd := newTileCmd("WHEREIN", strconv.Itoa(len(values)))
	for _, val := range values {
		cmd.appendArgs(floatString(val))
	}

	query.searchOpts = append(query.searchOpts, cmd)
	return query
}

// Match is similar to WHERE except that it works on the object id instead of fields.
// There can be multiple MATCH options in a single search.
// The MATCH value is a simple glob pattern.
func (query GeofenceQueryBuilder) Match(pattern string) GeofenceQueryBuilder {
	query.searchOpts = append(query.searchOpts, newTileCmd("MATCH", pattern))
	return query
}

// Format set response format.
func (query GeofenceQueryBuilder) Format(fmt OutputFormat) GeofenceQueryBuilder {
	query.outputFormat = fmt
	return query
}
