package t38c

import "context"

// ScanQueryBuilder struct
type ScanQueryBuilder struct {
	client       tile38Client
	key          string
	outputFormat *OutputFormat
	opts         searchOpts
}

func newScanQueryBuilder(client tile38Client, key string) ScanQueryBuilder {
	return ScanQueryBuilder{
		client: client,
		key:    key,
	}
}

func (query ScanQueryBuilder) toCmd() cmd {
	args := []string{query.key}
	args = append(args, query.opts.Args()...)
	if query.outputFormat != nil {
		args = append(args, query.outputFormat.Name)
		args = append(args, query.outputFormat.Args...)
	}

	return newCmd("SCAN", args...)
}

// Do cmd
func (query ScanQueryBuilder) Do(ctx context.Context) (*SearchResponse, error) {
	cmd := query.toCmd()
	resp := &SearchResponse{}
	err := query.client.jExecute(ctx, &resp, cmd.Name, cmd.Args...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Cursor is used to iterate though many objects from the search results.
// An iteration begins when the CURSOR is set to Zero or not included with the request,
// and completes when the cursor returned by the server is Zero.
func (query ScanQueryBuilder) Cursor(cursor int) ScanQueryBuilder {
	query.opts.Cursor = &cursor
	return query
}

// Limit can be used to limit the number of objects returned for a single search request.
func (query ScanQueryBuilder) Limit(limit int) ScanQueryBuilder {
	query.opts.Limit = &limit
	return query
}

// Match is similar to WHERE except that it works on the object id instead of fields.
// There can be multiple MATCH options in a single search.
// The MATCH value is a simple glob pattern.
func (query ScanQueryBuilder) Match(pattern string) ScanQueryBuilder {
	query.opts.Match = append(query.opts.Match, pattern)
	return query
}

// Asc order. Only for SEARCH and SCAN commands.
func (query ScanQueryBuilder) Asc() ScanQueryBuilder {
	query.opts.Asc = true
	return query
}

// Desc order. Only for SEARCH and SCAN commands.
func (query ScanQueryBuilder) Desc() ScanQueryBuilder {
	query.opts.Desc = true
	return query
}

// Where allows for filtering out results based on field values.
func (query ScanQueryBuilder) Where(field string, min, max float64) ScanQueryBuilder {
	query.opts.Where = append(query.opts.Where, whereOpt{
		Field: field,
		Min:   min,
		Max:   max,
	})
	return query
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func (query ScanQueryBuilder) Wherein(field string, values ...float64) ScanQueryBuilder {
	query.opts.Wherein = append(query.opts.Wherein, whereinOpt{
		Field:  field,
		Values: values,
	})
	return query
}

// NoFields tells the server that you do not want field values returned with the search results.
func (query ScanQueryBuilder) NoFields() ScanQueryBuilder {
	query.opts.NoFields = true
	return query
}

// Format set response format.
func (query ScanQueryBuilder) Format(fmt OutputFormat) ScanQueryBuilder {
	query.outputFormat = &fmt
	return query
}

// RawQuery is similar to a Where clause, but it allows for the input of a raw query.
// As of Tile38 1.30.0 FIELD is no longer limited to numbers.
// Example can be searched for with `SCAN fleet WHERE driver.firstname == Josh IDS`
// or `INTERSECTS fleet WHERE 'info.speed > 45 && info.age < 21' BOUNDS 30 -120 40 -100`
func (query ScanQueryBuilder) RawQuery(rawQuery string) ScanQueryBuilder {
	query.opts.RawQuery = append(query.opts.RawQuery, rawQuery)
	return query
}
