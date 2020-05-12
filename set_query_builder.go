package t38c

import "strconv"

// SetQueryBuilder struct
type SetQueryBuilder struct {
	client     *Client
	key        string
	objectID   string
	area       Command
	fields     []Field
	nx         bool
	xx         bool
	expiration *int
}

func newSetQueryBuilder(client *Client, key, objectID string, area Command) SetQueryBuilder {
	return SetQueryBuilder{
		client:   client,
		key:      key,
		objectID: objectID,
		area:     area,
	}
}

func (query SetQueryBuilder) toCmd() Command {
	args := []string{
		query.key, query.objectID,
	}

	if query.nx {
		args = append(args, "NX")
	}

	if query.xx {
		args = append(args, "XX")
	}

	if query.expiration != nil {
		args = append(args, "EX")
		args = append(args, strconv.Itoa(*query.expiration))
	}

	for _, field := range query.fields {
		args = append(args, "FIELD")
		args = append(args, field.Name)
		args = append(args, floatString(field.Value))
	}

	args = append(args, query.area.Name)
	args = append(args, query.area.Args...)

	return NewCommand("SET", args...)
}

// Do cmd
func (query SetQueryBuilder) Do() error {
	cmd := query.toCmd()
	return query.client.jExecute(nil, cmd.Name, cmd.Args...)
}

// Field sets the object field
func (query SetQueryBuilder) Field(name string, value float64) SetQueryBuilder {
	query.fields = append(query.fields, Field{name, value})
	return query
}

// IfNotExists only set the object if it does not already exist
func (query SetQueryBuilder) IfNotExists() SetQueryBuilder {
	query.nx = true
	return query
}

// IfExists only set the object if it already exist
func (query SetQueryBuilder) IfExists() SetQueryBuilder {
	query.xx = true
	return query
}

// Expiration sets the specified expire time, in seconds
func (query SetQueryBuilder) Expiration(seconds int) SetQueryBuilder {
	query.expiration = &seconds
	return query
}
