package t38c

import "strconv"

type WhereOpt struct {
	Field string
	Min   float64
	Max   float64
}

type WhereinOpt struct {
	Field  string
	Values []float64
}

type SearchQueryParams struct {
	Asc      bool
	Desc     bool
	NoFields bool
	Clip     bool
	Distance bool
	Cursor   *int
	Limit    *int
	Sparse   *int
	Where    *WhereOpt
	Wherein  *WhereinOpt
	Match    *string
}

type SearchQuery struct {
	Cmd          string
	Key          string
	Area         Command
	OutputFormat *OutputFormat
	Params       SearchQueryParams
}

type searchQueryBuilder struct {
	cmd          string
	key          string
	area         Command
	outputFormat OutputFormat
	opts         []Command
}

// Cmd builds tile38 search command.
func (query searchQueryBuilder) Cmd() Command {
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

// Asc order. Only for SEARCH and SCAN commands.
func (query searchQueryBuilder) Asc() searchQueryBuilder {
	query.opts = append(query.opts, NewCommand("ASC"))
	return query
}

// Desc order. Only for SEARCH and SCAN commands.
func (query searchQueryBuilder) Desc() searchQueryBuilder {
	query.opts = append(query.opts, NewCommand("DESC"))
	return query
}

// NoFields tells the server that you do not want field values returned with the search results.
func (query searchQueryBuilder) NoFields() searchQueryBuilder {
	query.opts = append(query.opts, NewCommand("NOFIELDS"))
	return query
}

// Clip tells the server to clip intersecting objects by the bounding box area of the search.
// It can only be used with these area formats: BOUNDS, TILE, QUADKEY, HASH.
func (query searchQueryBuilder) Clip() searchQueryBuilder {
	query.opts = append(query.opts, NewCommand("CLIP"))
	return query
}

// Distance allows to return between objects. Only for NEARBY command.
func (query searchQueryBuilder) Distance() searchQueryBuilder {
	query.opts = append(query.opts, NewCommand("DISTANCE"))
	return query
}

// Cursor is used to iterate though many objects from the search results.
// An iteration begins when the CURSOR is set to Zero or not included with the request,
// and completes when the cursor returned by the server is Zero.
func (query searchQueryBuilder) Cursor(cursor int) searchQueryBuilder {
	query.opts = append(query.opts, NewCommand("CURSOR", strconv.Itoa(cursor)))
	return query
}

// Limit can be used to limit the number of objects returned for a single search request.
func (query searchQueryBuilder) Limit(limit int) searchQueryBuilder {
	query.opts = append(query.opts, NewCommand("LIMIT", strconv.Itoa(limit)))
	return query
}

// Sparse will distribute the results of a search evenly across the requested area.
func (query searchQueryBuilder) Sparse(sparse int) searchQueryBuilder {
	query.opts = append(query.opts, NewCommand("SPARSE", strconv.Itoa(sparse)))
	return query
}

// Where allows for filtering out results based on field values.
func (query searchQueryBuilder) Where(field string, min, max float64) searchQueryBuilder {
	query.opts = append(query.opts, NewCommand("WHERE", field, floatString(min), floatString(max)))
	return query
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func (query searchQueryBuilder) Wherein(field string, values ...float64) searchQueryBuilder {
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
func (query searchQueryBuilder) Match(pattern string) searchQueryBuilder {
	query.opts = append(query.opts, NewCommand("MATCH", pattern))
	return query
}

// Format set response format.
func (query searchQueryBuilder) Format(fmt OutputFormat) searchQueryBuilder {
	query.outputFormat = fmt
	return query
}

// SetOption ...
type SetOption Command

var (
	// IfNotExists only set the object if it does not already exist.
	IfNotExists = SetOption(NewCommand("NX"))

	// IfExists only set the object if it already exist.
	IfExists = SetOption(NewCommand("XX"))
)

// Field are extra data which belongs to an object.
// A field is always a double precision floating point.
// There is no limit to the number of fields that an object can have.
func Field(name string, value float64) SetOption {
	return SetOption(NewCommand("FIELD", name, floatString(value)))
}

// Expiration set the specified expire time, in seconds.
func Expiration(seconds int) SetOption {
	return SetOption(NewCommand("EX", strconv.Itoa(seconds)))
}
