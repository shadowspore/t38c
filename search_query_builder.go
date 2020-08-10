package t38c

// SearchQueryBuilder struct
type SearchQueryBuilder struct {
	client       tile38Client
	key          string
	outputFormat *OutputFormat
	opts         searchOpts
}

func newSearchQueryBuilder(client tile38Client, key string) SearchQueryBuilder {
	return SearchQueryBuilder{
		client: client,
		key:    key,
	}
}

func (query SearchQueryBuilder) toCmd() cmd {
	args := []string{query.key}

	if query.outputFormat != nil {
		args = append(args, query.outputFormat.Name)
		args = append(args, query.outputFormat.Args...)
	}

	return newCmd("SEARCH", args...)
}

// Do cmd
func (query SearchQueryBuilder) Do() (resp *SearchResponse, err error) {
	cmd := query.toCmd()
	err = query.client.jExecute(&resp, cmd.Name, cmd.Args...)
	return
}

// Cursor is used to iterate though many objects from the search results.
// An iteration begins when the CURSOR is set to Zero or not included with the request,
// and completes when the cursor returned by the server is Zero.
func (query SearchQueryBuilder) Cursor(cursor int) SearchQueryBuilder {
	query.opts.Curosr = &cursor
	return query
}

// Limit can be used to limit the number of objects returned for a single search request.
func (query SearchQueryBuilder) Limit(limit int) SearchQueryBuilder {
	query.opts.Limit = &limit
	return query
}

// Match is similar to WHERE except that it works on the object id instead of fields.
// There can be multiple MATCH options in a single search.
// The MATCH value is a simple glob pattern.
func (query SearchQueryBuilder) Match(pattern string) SearchQueryBuilder {
	query.opts.Match = append(query.opts.Match, pattern)
	return query
}

// Asc order. Only for SEARCH and SCAN commands.
func (query SearchQueryBuilder) Asc() SearchQueryBuilder {
	query.opts.Asc = true
	return query
}

// Desc order. Only for SEARCH and SCAN commands.
func (query SearchQueryBuilder) Desc() SearchQueryBuilder {
	query.opts.Desc = true
	return query
}

// Where allows for filtering out results based on field values.
func (query SearchQueryBuilder) Where(field string, min, max float64) SearchQueryBuilder {
	query.opts.Where = append(query.opts.Where, whereOpt{
		Field: field,
		Min:   min,
		Max:   max,
	})
	return query
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func (query SearchQueryBuilder) Wherein(field string, values ...float64) SearchQueryBuilder {
	query.opts.Wherein = append(query.opts.Wherein, whereinOpt{
		Field:  field,
		Values: values,
	})
	return query
}

// NoFields tells the server that you do not want field values returned with the search results.
func (query SearchQueryBuilder) NoFields() SearchQueryBuilder {
	query.opts.NoFields = true
	return query
}

// FormatCount - total object count sent in the response.
func (query SearchQueryBuilder) FormatCount() SearchQueryBuilder {
	query.outputFormat = &FormatCount
	return query
}

// FormatIDs - a list of IDs belonging to the key. Will not return the objects.
func (query SearchQueryBuilder) FormatIDs() SearchQueryBuilder {
	query.outputFormat = &FormatIDs
	return query
}
