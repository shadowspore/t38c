package t38c

import (
	"context"
	"strconv"
)

// SetQueryBuilder struct
type SetQueryBuilder struct {
	client     tile38Client
	key        string
	objectID   string
	area       cmd
	fields     []field
	nx         bool
	xx         bool
	expiration *int
}

func newSetQueryBuilder(client tile38Client, key, objectID string, area cmd) SetQueryBuilder {
	return SetQueryBuilder{
		client:   client,
		key:      key,
		objectID: objectID,
		area:     area,
	}
}

func (query SetQueryBuilder) toCmd() cmd {
	args := []string{query.key, query.objectID}
	if query.nx {
		args = append(args, "NX")
	}

	if query.xx {
		args = append(args, "XX")
	}

	if query.expiration != nil {
		args = append(args, "EX", strconv.Itoa(*query.expiration))
	}

	for _, field := range query.fields {
		args = append(args, "FIELD", field.Name, floatString(field.Value))
	}

	args = append(args, query.area.Name)
	args = append(args, query.area.Args...)
	return newCmd("SET", args...)
}

// Do cmd
func (query SetQueryBuilder) Do(ctx context.Context) error {
	cmd := query.toCmd()
	return query.client.jExecute(ctx, nil, cmd.Name, cmd.Args...)
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
