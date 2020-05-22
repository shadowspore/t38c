package t38c

import "strconv"

// SearchQueryBuilder struct
type SearchQueryBuilder struct {
	client       *Client
	key          string
	outputFormat OutputFormat
	opts         []Command
}

func newSearchQueryBuilder(client *Client, key string) SearchQueryBuilder {
	return SearchQueryBuilder{
		client: client,
		key:    key,
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

	return NewCommand("SEARCH", args...)
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

// Match is similar to WHERE except that it works on the object id instead of fields.
// There can be multiple MATCH options in a single search.
// The MATCH value is a simple glob pattern.
func (query SearchQueryBuilder) Match(pattern string) SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("MATCH", pattern))
	return query
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

// NoFields tells the server that you do not want field values returned with the search results.
func (query SearchQueryBuilder) NoFields() SearchQueryBuilder {
	query.opts = append(query.opts, NewCommand("NOFIELDS"))
	return query
}

// FormatCount - total object count sent in the response.
func (query SearchQueryBuilder) FormatCount() SearchQueryBuilder {
	query.outputFormat = FormatCount
	return query
}

// FormatIDs - a list of IDs belonging to the key. Will not return the objects.
func (query SearchQueryBuilder) FormatIDs() SearchQueryBuilder {
	query.outputFormat = FormatIDs
	return query
}
