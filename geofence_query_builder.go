package t38c

import (
	"context"
	"strconv"
	"strings"
)

// GeofenceQueryBuilder optional params
type GeofenceQueryBuilder struct {
	client         tile38Client
	isRoamQuery    bool
	cmd            string
	key            string
	area           cmd
	target         string
	pattern        string
	meters         int
	outputFormat   *OutputFormat
	detectActions  []DetectAction
	notifyCommands []NotifyCommand
	searchOpts     searchOpts
	isNodwell      bool
}

func newGeofenceQueryBuilder(client tile38Client, cmd, key string, area cmd) GeofenceQueryBuilder {
	return GeofenceQueryBuilder{
		client: client,
		cmd:    cmd,
		key:    key,
		area:   area,
	}
}

func newGeofenceRoamQueryBuilder(client tile38Client, key, target, pattern string, meters int, nodwell bool) GeofenceQueryBuilder {
	return GeofenceQueryBuilder{
		client:      client,
		cmd:         "NEARBY",
		isRoamQuery: true,
		key:         key,
		target:      target,
		pattern:     pattern,
		meters:      meters,
		isNodwell:   nodwell,
	}
}

func (query GeofenceQueryBuilder) toCmd() cmd {
	args := []string{query.key}
	args = append(args, query.searchOpts.Args()...)
	args = append(args, "FENCE")
	if query.isNodwell {
		args = append(args, "NODWELL")
	}

	if len(query.detectActions) > 0 {
		actions := make([]string, len(query.detectActions))
		for i := range query.detectActions {
			actions[i] = string(query.detectActions[i])
		}
		args = append(args, "DETECT", strings.Join(actions, ","))
	}

	if len(query.notifyCommands) > 0 {
		commands := make([]string, len(query.notifyCommands))
		for i := range query.notifyCommands {
			commands[i] = string(query.notifyCommands[i])
		}
		args = append(args, "COMMANDS", strings.Join(commands, ","))
	}

	if query.outputFormat != nil {
		args = append(args, query.outputFormat.Name)
		args = append(args, query.outputFormat.Args...)
	}

	if query.isRoamQuery {
		args = append(args, "ROAM", query.target, query.pattern, strconv.Itoa(query.meters))
	} else {
		args = append(args, query.area.Name)
		args = append(args, query.area.Args...)
	}

	return newCmd(query.cmd, args...)
}

// Do cmd
func (query GeofenceQueryBuilder) Do(ctx context.Context, handler EventHandler) error {
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

// NoFields tells the server that you do not want field values returned with the search results.
func (query GeofenceQueryBuilder) NoFields() GeofenceQueryBuilder {
	query.searchOpts.NoFields = true
	return query
}

// Clip tells the server to clip intersecting objects by the bounding box area of the search.
// It can only be used with these area formats: BOUNDS, TILE, QUADKEY, HASH.
func (query GeofenceQueryBuilder) Clip() GeofenceQueryBuilder {
	query.searchOpts.Clip = true
	return query
}

// Distance allows to return between objects. Only for NEARBY tileCmd.
func (query GeofenceQueryBuilder) Distance() GeofenceQueryBuilder {
	query.searchOpts.Distance = true
	return query
}

// Cursor is used to iterate though many objects from the search results.
// An iteration begins when the CURSOR is set to Zero or not included with the request,
// and completes when the cursor returned by the server is Zero.
func (query GeofenceQueryBuilder) Cursor(cursor int) GeofenceQueryBuilder {
	query.searchOpts.Cursor = &cursor
	return query
}

// Limit can be used to limit the number of objects returned for a single search request.
func (query GeofenceQueryBuilder) Limit(limit int) GeofenceQueryBuilder {
	query.searchOpts.Limit = &limit
	return query
}

// Sparse will distribute the results of a search evenly across the requested area.
func (query GeofenceQueryBuilder) Sparse(sparse int) GeofenceQueryBuilder {
	query.searchOpts.Sparse = &sparse
	return query
}

// Where allows for filtering out results based on field values.
func (query GeofenceQueryBuilder) Where(field string, min, max float64) GeofenceQueryBuilder {
	query.searchOpts.Where = append(query.searchOpts.Where, whereOpt{
		Field: field,
		Min:   min,
		Max:   max,
	})
	return query
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func (query GeofenceQueryBuilder) Wherein(field string, values ...float64) GeofenceQueryBuilder {
	query.searchOpts.Wherein = append(query.searchOpts.Wherein, whereinOpt{
		Field:  field,
		Values: values,
	})
	return query
}

// WhereEval similar to WHERE except that matching decision is made by Lua script
// For example:
// 'nearby fleet whereeval "return FIELDS.wheels > ARGV[1] or (FIELDS.length * FIELDS.width) > ARGV[2]" 2 8 120 point 33.462 -112.268 6000'
// will return only the objects in the fleet collection that are within the 6km radius
// and have a field named wheels that is above 8, or have length and width whose product is greater than 120.
// Multiple WHEREEVALs are concatenated as and clauses. See EVAL command for more details.
// Note that, unlike the EVAL command, WHEREVAL Lua environment (1) does not have KEYS global,
// and (2) has the FIELDS global with the Lua table of the iterated object’s fields.
func (query GeofenceQueryBuilder) WhereEval(script string, args ...string) GeofenceQueryBuilder {
	query.searchOpts.WhereEval = append(query.searchOpts.WhereEval, whereEvalOpt{
		Name: script,
		Args: args,
	})
	return query
}

// WhereEvalSHA similar to WHERE except that matching decision is made by Lua script
// For example:
// 'nearby fleet whereeval "return FIELDS.wheels > ARGV[1] or (FIELDS.length * FIELDS.width) > ARGV[2]" 2 8 120 point 33.462 -112.268 6000'
// will return only the objects in the fleet collection that are within the 6km radius
// and have a field named wheels that is above 8, or have length and width whose product is greater than 120.
// Multiple WHEREEVALs are concatenated as and clauses. See EVAL command for more details.
// Note that, unlike the EVAL command, WHEREVAL Lua environment (1) does not have KEYS global,
// and (2) has the FIELDS global with the Lua table of the iterated object’s fields.
func (query GeofenceQueryBuilder) WhereEvalSHA(sha string, args ...string) GeofenceQueryBuilder {
	query.searchOpts.WhereEval = append(query.searchOpts.WhereEval, whereEvalOpt{
		Name:  sha,
		IsSHA: true,
		Args:  args,
	})
	return query
}

// Match is similar to WHERE except that it works on the object id instead of fields.
// There can be multiple MATCH options in a single search.
// The MATCH value is a simple glob pattern.
func (query GeofenceQueryBuilder) Match(pattern string) GeofenceQueryBuilder {
	query.searchOpts.Match = append(query.searchOpts.Match, pattern)
	return query
}

// Format set response format.
func (query GeofenceQueryBuilder) Format(fmt OutputFormat) GeofenceQueryBuilder {
	query.outputFormat = &fmt
	return query
}

// RawQuery is similar to a Where clause, but it allows for the input of a raw query.
// As of Tile38 1.30.0 FIELD is no longer limited to numbers.
// Example can be searched for with `SCAN fleet WHERE driver.firstname == Josh IDS`
// or `INTERSECTS fleet WHERE 'info.speed > 45 && info.age < 21' BOUNDS 30 -120 40 -100`
func (query GeofenceQueryBuilder) RawQuery(rawQuery string) GeofenceQueryBuilder {
	query.searchOpts.RawQuery = append(query.searchOpts.RawQuery, rawQuery)
	return query
}
