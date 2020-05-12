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

// SearchQueryBuilder struct
type SearchQueryBuilder struct {
	client       *Client
	cmd          string
	key          string
	area         Command
	outputFormat OutputFormat
	opts         []Command
}

func newSearchQueryBuilder(client *Client, cmd, key string, area Command) SearchQueryBuilder {
	return SearchQueryBuilder{
		client: client,
		cmd:    cmd,
		key:    key,
		area:   area,
	}
}

func (query SearchQueryBuilder) toCmd() Command {
	var args []string
	args = append(args, query.key)

	for _, opt := range query.opts {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}

	if len(query.outputFormat.Name) > 0 {
		args = append(args, query.outputFormat.Name)
		args = append(args, query.outputFormat.Args...)
	}

	args = append(args, query.area.Name)
	args = append(args, query.area.Args...)

	return NewCommand(query.cmd, args...)
}

// Do cmd
func (query SearchQueryBuilder) Do() (*SearchResponse, error) {
	cmd := query.toCmd()
	resp := &SearchResponse{}
	err := query.client.jExecute(&resp, cmd.Name, cmd.Args...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Asc order. Only for SEARCH and SCAN commands.
func (query SearchQueryBuilder) Asc() SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("ASC"))
	return query
}

// Desc order. Only for SEARCH and SCAN commands.
func (query SearchQueryBuilder) Desc() SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("DESC"))
	return query
}

// NoFields tells the server that you do not want field values returned with the search results.
func (query SearchQueryBuilder) NoFields() SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("NOFIELDS"))
	return query
}

// Clip tells the server to clip intersecting objects by the bounding box area of the search.
// It can only be used with these area formats: BOUNDS, TILE, QUADKEY, HASH.
func (query SearchQueryBuilder) Clip() SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("CLIP"))
	return query
}

// Distance allows to return between objects. Only for NEARBY command.
func (query SearchQueryBuilder) Distance() SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("DISTANCE"))
	return query
}

// Cursor is used to iterate though many objects from the search results.
// An iteration begins when the CURSOR is set to Zero or not included with the request,
// and completes when the cursor returned by the server is Zero.
func (query SearchQueryBuilder) Cursor(cursor int) SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("CURSOR", strconv.Itoa(cursor)))
	return query
}

// Limit can be used to limit the number of objects returned for a single search request.
func (query SearchQueryBuilder) Limit(limit int) SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("LIMIT", strconv.Itoa(limit)))
	return query
}

// Sparse will distribute the results of a search evenly across the requested area.
func (query SearchQueryBuilder) Sparse(sparse int) SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("SPARSE", strconv.Itoa(sparse)))
	return query
}

// Where allows for filtering out results based on field values.
func (query SearchQueryBuilder) Where(field string, min, max float64) SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("WHERE", field, floatString(min), floatString(max)))
	return query
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func (query SearchQueryBuilder) Wherein(field string, values ...float64) SearchQueryBuilder {
	var args []string
	args = append(args, strconv.Itoa(len(values)))
	for _, val := range values {
		args = append(args, floatString(val))
	}

	query.opts = append(query.opts, NewCommand("WHEREIN", args...))
	return query
}

// Match is similar to WHERE except that it works on the object id instead of fields.
// There can be multiple MATCH options in a single search.
// The MATCH value is a simple glob pattern.
func (query SearchQueryBuilder) Match(pattern string) SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("MATCH", pattern))
	return query
}

// Format set response format.
func (query SearchQueryBuilder) Format(fmt OutputFormat) SearchQueryBuilder {
	query.outputFormat = fmt
	return query
}
