package t38c

// FSetQueryBuilder struct
type FSetQueryBuilder struct {
	client   *Client
	key      string
	objectID string
	fields   []Field
	xx       bool
}

func (query FSetQueryBuilder) toCmd() Command {
	var args []string = []string{
		query.key, query.objectID,
	}

	if query.xx {
		args = append(args, "XX")
	}

	for _, field := range query.fields {
		args = append(args, field.Name)
		args = append(args, floatString(field.Value))
	}

	return NewCommand("FSET", args...)
}

// Do cmd
func (query FSetQueryBuilder) Do() error {
	cmd := query.toCmd()
	return query.client.jExecute(nil, cmd.Name, cmd.Args...)
}

// Field sets the object field
func (query FSetQueryBuilder) Field(name string, value float64) FSetQueryBuilder {
	query.fields = append(query.fields, Field{name, value})
	return query
}

// IfExists only set the object if it already exist
func (query FSetQueryBuilder) IfExists() FSetQueryBuilder {
	query.xx = true
	return query
}
