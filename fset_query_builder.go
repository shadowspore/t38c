package t38c

// FSetQueryBuilder struct
type FSetQueryBuilder struct {
	client   *Client
	key      string
	objectID string
	fields   []field
	xx       bool
}

func newFSetQueryBuilder(client *Client, key, objectID string) FSetQueryBuilder {
	return FSetQueryBuilder{
		client:   client,
		key:      key,
		objectID: objectID,
	}
}

func (query FSetQueryBuilder) toCmd() *tileCmd {
	cmd := newTileCmd("FSET", query.key, query.objectID)
	if query.xx {
		cmd.appendArgs("XX")
	}

	for _, field := range query.fields {
		cmd.appendArgs(field.Name, floatString(field.Value))
	}

	return cmd
}

// Do cmd
func (query FSetQueryBuilder) Do() error {
	cmd := query.toCmd()
	return query.client.jExecute(nil, cmd.Name, cmd.Args...)
}

// Field sets the object field
func (query FSetQueryBuilder) Field(name string, value float64) FSetQueryBuilder {
	query.fields = append(query.fields, field{name, value})
	return query
}

// IfExists only set the object if it already exist
func (query FSetQueryBuilder) IfExists() FSetQueryBuilder {
	query.xx = true
	return query
}
