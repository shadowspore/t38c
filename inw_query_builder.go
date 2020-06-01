package t38c

import "strconv"

// type WhereOpt struct {
// 	Field string
// 	Min   float64
// 	Max   float64
// }

// type WhereinOpt struct {
// 	Field  string
// 	Values []float64
// }

// type SearchQueryParams struct {
// 	Asc      bool
// 	Desc     bool
// 	NoFields bool
// 	Clip     bool
// 	Distance bool
// 	Cursor   *int
// 	Limit    *int
// 	Sparse   *int
// 	Where    []*WhereOpt
// 	Wherein  []*WhereinOpt
// 	Match    []*string
// }

// type SearchQuery struct {
// 	Cmd          string
// 	Key          string
// 	Area         Command
// 	OutputFormat *OutputFormat
// 	Params       SearchQueryParams
// }

// InwQueryBuilder struct
// Intersects Nearby Within
type InwQueryBuilder struct {
	client       *Client
	cmd          string
	key          string
	area         tileCmd
	outputFormat OutputFormat
	opts         []tileCmd
}

func newInwQueryBuilder(client *Client, cmd, key string, area tileCmd) InwQueryBuilder {
	return InwQueryBuilder{
		client: client,
		cmd:    cmd,
		key:    key,
		area:   area,
	}
}

func (query InwQueryBuilder) toCmd() tileCmd {
	cmd := newTileCmd(query.cmd, query.key)
	for _, opt := range query.opts {
		cmd.appendArgs(opt.Name, opt.Args...)
	}

	if len(query.outputFormat.Name) > 0 {
		cmd.appendArgs(query.outputFormat.Name, query.outputFormat.Args...)
	}

	if len(query.area.Name) > 0 {
		cmd.appendArgs(query.area.Name, query.area.Args...)
	}

	return cmd
}

// Do cmd
func (query InwQueryBuilder) Do() (*SearchResponse, error) {
	cmd := query.toCmd()
	resp := &SearchResponse{}
	err := query.client.jExecute(&resp, cmd.Name, cmd.Args...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Cursor is used to iterate though many objects from the search results.
// An iteration begins when the CURSOR is set to Zero or not included with the request,
// and completes when the cursor returned by the server is Zero.
func (query InwQueryBuilder) Cursor(cursor int) InwQueryBuilder {
	query.opts = append(query.opts, newTileCmd("CURSOR", strconv.Itoa(cursor)))
	return query
}

// Limit can be used to limit the number of objects returned for a single search request.
func (query InwQueryBuilder) Limit(limit int) InwQueryBuilder {
	query.opts = append(query.opts, newTileCmd("LIMIT", strconv.Itoa(limit)))
	return query
}

// Sparse will distribute the results of a search evenly across the requested area.
func (query InwQueryBuilder) Sparse(sparse int) InwQueryBuilder {
	query.opts = append(query.opts, newTileCmd("SPARSE", strconv.Itoa(sparse)))
	return query
}

// Match is similar to WHERE except that it works on the object id instead of fields.
// There can be multiple MATCH options in a single search.
// The MATCH value is a simple glob pattern.
func (query InwQueryBuilder) Match(pattern string) InwQueryBuilder {
	query.opts = append(query.opts, newTileCmd("MATCH", pattern))
	return query
}

// Distance allows to return between objects.
// Only for NEARBY command.
func (query InwQueryBuilder) Distance() InwQueryBuilder {
	query.opts = append(query.opts, newTileCmd("DISTANCE"))
	return query
}

// Where allows for filtering out results based on field values.
func (query InwQueryBuilder) Where(field string, min, max float64) InwQueryBuilder {
	query.opts = append(query.opts, newTileCmd("WHERE", field, floatString(min), floatString(max)))
	return query
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func (query InwQueryBuilder) Wherein(field string, values ...float64) InwQueryBuilder {
	var args []string
	args = append(args, strconv.Itoa(len(values)))
	for _, val := range values {
		args = append(args, floatString(val))
	}

	query.opts = append(query.opts, newTileCmd("WHEREIN", args...))
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
	cmd := newTileCmd("WHEREEVAL", append([]string{script}, args...)...)
	query.opts = append(query.opts, cmd)
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
	cmd := newTileCmd("WHEREEVALSHA", append([]string{sha}, args...)...)
	query.opts = append(query.opts, cmd)
	return query
}

// Clip tells the server to clip intersecting objects by the bounding box area of the search.
// It can only be used with these area formats: BOUNDS, TILE, QUADKEY, HASH.
// Only for INTERSECTS command.
func (query InwQueryBuilder) Clip() InwQueryBuilder {
	query.opts = append(query.opts, newTileCmd("CLIP"))
	return query
}

// NoFields tells the server that you do not want field values returned with the search results.
func (query InwQueryBuilder) NoFields() InwQueryBuilder {
	query.opts = append(query.opts, newTileCmd("NOFIELDS"))
	return query
}

// Format set response format.
func (query InwQueryBuilder) Format(fmt OutputFormat) InwQueryBuilder {
	query.outputFormat = fmt
	return query
}
