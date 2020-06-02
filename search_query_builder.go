package t38c

import "strconv"

// SearchQueryBuilder struct
type SearchQueryBuilder struct {
	client       *Client
	key          string
	outputFormat OutputFormat
	opts         []*tileCmd
}

func newSearchQueryBuilder(client *Client, key string) SearchQueryBuilder {
	return SearchQueryBuilder{
		client: client,
		key:    key,
	}
}

func (query SearchQueryBuilder) toCmd() *tileCmd {
	cmd := newTileCmd("SEARCH", query.key)
	for _, opt := range query.opts {
		cmd.appendArgs(opt.Name, opt.Args...)
	}

	if query.outputFormat != nil {
		cmd.appendArgs(query.outputFormat.Name, query.outputFormat.Args...)
	}

	return cmd
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
	query.opts = append(query.opts, newTileCmd("CURSOR", strconv.Itoa(cursor)))
	return query
}

// Limit can be used to limit the number of objects returned for a single search request.
func (query SearchQueryBuilder) Limit(limit int) SearchQueryBuilder {
	query.opts = append(query.opts, newTileCmd("LIMIT", strconv.Itoa(limit)))
	return query
}

// Match is similar to WHERE except that it works on the object id instead of fields.
// There can be multiple MATCH options in a single search.
// The MATCH value is a simple glob pattern.
func (query SearchQueryBuilder) Match(pattern string) SearchQueryBuilder {
	query.opts = append(query.opts, newTileCmd("MATCH", pattern))
	return query
}

// Asc order. Only for SEARCH and SCAN commands.
func (query SearchQueryBuilder) Asc() SearchQueryBuilder {
	query.opts = append(query.opts, newTileCmd("ASC"))
	return query
}

// Desc order. Only for SEARCH and SCAN commands.
func (query SearchQueryBuilder) Desc() SearchQueryBuilder {
	query.opts = append(query.opts, newTileCmd("DESC"))
	return query
}

// Where allows for filtering out results based on field values.
func (query SearchQueryBuilder) Where(field string, min, max float64) SearchQueryBuilder {
	query.opts = append(query.opts, newTileCmd("WHERE", field, floatString(min), floatString(max)))
	return query
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func (query SearchQueryBuilder) Wherein(field string, values ...float64) SearchQueryBuilder {
	cmd := newTileCmd("WHEREIN", field, strconv.Itoa(len(values)))
	for _, val := range values {
		cmd.appendArgs(floatString(val))
	}

	query.opts = append(query.opts, cmd)
	return query
}

// NoFields tells the server that you do not want field values returned with the search results.
func (query SearchQueryBuilder) NoFields() SearchQueryBuilder {
	query.opts = append(query.opts, newTileCmd("NOFIELDS"))
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
