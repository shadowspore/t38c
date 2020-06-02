package t38c

import "strconv"

// SetQueryBuilder struct
type SetQueryBuilder struct {
	client     tile38Client
	key        string
	objectID   string
	area       *tileCmd
	fields     []field
	nx         bool
	xx         bool
	expiration *int
}

func newSetQueryBuilder(client tile38Client, key, objectID string, area *tileCmd) SetQueryBuilder {
	return SetQueryBuilder{
		client:   client,
		key:      key,
		objectID: objectID,
		area:     area,
	}
}

func (query SetQueryBuilder) toCmd() *tileCmd {
	cmd := newTileCmd("SET", query.key, query.objectID)
	if query.nx {
		cmd.appendArgs("NX")
	}

	if query.xx {
		cmd.appendArgs("XX")
	}

	if query.expiration != nil {
		cmd.appendArgs("EX", strconv.Itoa(*query.expiration))
	}

	for _, field := range query.fields {
		cmd.appendArgs("FIELD", field.Name, floatString(field.Value))
	}

	cmd.appendArgs(query.area.Name, query.area.Args...)
	return cmd
}

// Do cmd
func (query SetQueryBuilder) Do() error {
	cmd := query.toCmd()
	return query.client.jExecute(nil, cmd.Name, cmd.Args...)
}

// Field sets the object field
func (query SetQueryBuilder) Field(name string, value float64) SetQueryBuilder {
	query.fields = append(query.fields, field{name, value})
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
