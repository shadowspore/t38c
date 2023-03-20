package t38c

import "context"

// InwQueryBuilder struct
// Intersects Nearby Within
type InwQueryBuilder struct {
	client       tile38Client
	cmd          string
	key          string
	area         cmd
	outputFormat *OutputFormat
	searchOpts   searchOpts
}

func newInwQueryBuilder(client tile38Client, cmd, key string, area cmd) InwQueryBuilder {
	return InwQueryBuilder{
		client: client,
		cmd:    cmd,
		key:    key,
		area:   area,
	}
}

func (query InwQueryBuilder) toCmd() cmd {
	args := []string{query.key}
	args = append(args, query.searchOpts.Args()...)
	if query.outputFormat != nil {
		args = append(args, query.outputFormat.Name)
		args = append(args, query.outputFormat.Args...)
	}

	args = append(args, query.area.Name)
	args = append(args, query.area.Args...)
	return newCmd(query.cmd, args...)
}

// Do cmd
func (query InwQueryBuilder) Do(ctx context.Context) (*SearchResponse, error) {
	cmd := query.toCmd()
	resp := new(SearchResponse)
	err := query.client.jExecute(ctx, resp, cmd.Name, cmd.Args...)
	return resp, err
}

// Cursor is used to iterate though many objects from the search results.
// An iteration begins when the CURSOR is set to Zero or not included with the request,
// and completes when the cursor returned by the server is Zero.
func (query InwQueryBuilder) Cursor(cursor int) InwQueryBuilder {
	query.searchOpts.Cursor = &cursor
	return query
}

// Limit can be used to limit the number of objects returned for a single search request.
func (query InwQueryBuilder) Limit(limit int) InwQueryBuilder {
	query.searchOpts.Limit = &limit
	return query
}

// Sparse will distribute the results of a search evenly across the requested area.
func (query InwQueryBuilder) Sparse(sparse int) InwQueryBuilder {
	query.searchOpts.Sparse = &sparse
	return query
}

// Match is similar to WHERE except that it works on the object id instead of fields.
// There can be multiple MATCH options in a single search.
// The MATCH value is a simple glob pattern.
func (query InwQueryBuilder) Match(pattern string) InwQueryBuilder {
	query.searchOpts.Match = append(query.searchOpts.Match, pattern)
	return query
}

// Distance allows to return between objects.
// Only for NEARBY command.
func (query InwQueryBuilder) Distance() InwQueryBuilder {
	query.searchOpts.Distance = true
	return query
}

// Where allows for filtering out results based on field values.
func (query InwQueryBuilder) Where(field string, min, max float64) InwQueryBuilder {
	query.searchOpts.Where = append(query.searchOpts.Where, whereOpt{
		Field: field,
		Min:   min,
		Max:   max,
	})
	return query
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func (query InwQueryBuilder) Wherein(field string, values ...float64) InwQueryBuilder {
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
func (query InwQueryBuilder) WhereEval(script string, args ...string) InwQueryBuilder {
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
func (query InwQueryBuilder) WhereEvalSHA(sha string, args ...string) InwQueryBuilder {
	query.searchOpts.WhereEval = append(query.searchOpts.WhereEval, whereEvalOpt{
		Name:  sha,
		IsSHA: true,
		Args:  args,
	})
	return query
}

// Clip tells the server to clip intersecting objects by the bounding box area of the search.
// It can only be used with these area formats: BOUNDS, TILE, QUADKEY, HASH.
// Only for INTERSECTS command.
func (query InwQueryBuilder) Clip() InwQueryBuilder {
	query.searchOpts.Clip = true
	return query
}

// NoFields tells the server that you do not want field values returned with the search results.
func (query InwQueryBuilder) NoFields() InwQueryBuilder {
	query.searchOpts.NoFields = true
	return query
}

// Format set response format.
func (query InwQueryBuilder) Format(fmt OutputFormat) InwQueryBuilder {
	query.outputFormat = &fmt
	return query
}

// RawQuery is similar to a Where clause, but it allows for the input of a raw query.
// As of Tile38 1.30.0 FIELD is no longer limited to numbers.
// Example can be searched for with `SCAN fleet WHERE driver.firstname == Josh IDS`
// or `INTERSECTS fleet WHERE 'info.speed > 45 && info.age < 21' BOUNDS 30 -120 40 -100`
func (query InwQueryBuilder) RawQuery(rawQuery string) InwQueryBuilder {
	query.searchOpts.RawQuery = append(query.searchOpts.RawQuery, rawQuery)
	return query
}
