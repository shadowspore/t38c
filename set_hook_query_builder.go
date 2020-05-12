package t38c

import (
	"strconv"
	"strings"
)

// SetHookQueryBuilder struct
type SetHookQueryBuilder struct {
	client     *Client
	name       string
	endpoints  []string
	cmd        Command
	metas      []Meta
	expiration *int
}

func newSetHookQueryBuilder(client *Client, name, endpoint string, cmd Command) SetHookQueryBuilder {
	return SetHookQueryBuilder{
		client:    client,
		name:      name,
		endpoints: []string{endpoint},
		cmd:       cmd,
	}
}

func (query SetHookQueryBuilder) toCmd() Command {
	args := []string{
		query.name,
		strings.Join(query.endpoints, ","),
	}

	for _, meta := range query.metas {
		args = append(args, "META")
		args = append(args, meta.Name)
		args = append(args, meta.Value)
	}

	if query.expiration != nil {
		args = append(args, "EX")
		args = append(args, strconv.Itoa(*query.expiration))
	}

	args = append(args, query.cmd.Name)
	args = append(args, query.cmd.Args...)
	return NewCommand("SETHOOK", args...)
}

// Do cmd
func (query SetHookQueryBuilder) Do() error {
	cmd := query.toCmd()
	return query.client.jExecute(nil, cmd.Name, cmd.Args...)
}

// Endpoint appends new endpoint to the hook.
// Tile38 will try to send a message to the first endpoint.
// If the send is a failure then the second endpoint is tried, and so on.
func (query SetHookQueryBuilder) Endpoint(endpoint string) SetHookQueryBuilder {
	query.endpoints = append(query.endpoints, endpoint)
	return query
}

// Expiration set the specified expire time, in seconds.
func (query SetHookQueryBuilder) Expiration(seconds int) SetHookQueryBuilder {
	query.expiration = &seconds
	return query
}

// Meta ...
func (query SetHookQueryBuilder) Meta(name, value string) SetHookQueryBuilder {
	query.metas = append(query.metas, Meta{name, value})
	return query
}
