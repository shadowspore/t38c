package t38c

import "strconv"

// SetChannelQueryBuilder struct
type SetChannelQueryBuilder struct {
	client     *Client
	name       string
	cmd        *tileCmd
	metas      []Meta
	expiration *int
}

func newSetChannelQueryBuilder(client *Client, name string, cmd *tileCmd) SetChannelQueryBuilder {
	return SetChannelQueryBuilder{
		client: client,
		name:   name,
		cmd:    cmd,
	}
}

func (query SetChannelQueryBuilder) toCmd() *tileCmd {
	cmd := newTileCmd("SETCHAN", query.name)
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
func (query SetChannelQueryBuilder) Do() error {
	cmd := query.toCmd()
	return query.client.jExecute(nil, cmd.Name, cmd.Args...)
}

// Expiration set the specified expire time, in seconds.
func (query SetChannelQueryBuilder) Expiration(seconds int) SetChannelQueryBuilder {
	query.expiration = &seconds
	return query
}

// Meta ...
func (query SetChannelQueryBuilder) Meta(name, value string) SetChannelQueryBuilder {
	query.metas = append(query.metas, Meta{name, value})
	return query
}
