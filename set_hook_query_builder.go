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
	cmd        *tileCmd
	metas      []Meta
	expiration *int
}

func newSetHookQueryBuilder(client *Client, name, endpoint string, cmd *tileCmd) SetHookQueryBuilder {
	return SetHookQueryBuilder{
		client:    client,
		name:      name,
		endpoints: []string{endpoint},
		cmd:       cmd,
	}
}

func (query SetHookQueryBuilder) toCmd() *tileCmd {
	cmd := newTileCmd("SETHOOK", query.name, strings.Join(query.endpoints, ","))
	for _, meta := range query.metas {
		cmd.appendArgs("META", meta.Name, meta.Value)
	}

	if query.expiration != nil {
		cmd.appendArgs("EX", strconv.Itoa(*query.expiration))
	}

	cmd.appendArgs(query.cmd.Name, query.cmd.Args...)
	return cmd
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
