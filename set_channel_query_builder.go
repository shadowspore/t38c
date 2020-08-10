package t38c

import "strconv"

// SetChannelQueryBuilder struct
type SetChannelQueryBuilder struct {
	client     tile38Client
	name       string
	cmd        cmd
	metas      []Meta
	expiration *int
}

func newSetChannelQueryBuilder(client tile38Client, name string, query cmd) SetChannelQueryBuilder {
	return SetChannelQueryBuilder{
		client: client,
		name:   name,
		cmd:    query,
	}
}

func (query SetChannelQueryBuilder) toCmd() cmd {
	args := []string{query.name}
	for _, meta := range query.metas {
		args = append(args, "META", meta.Name, meta.Value)
	}

	if query.expiration != nil {
		args = append(args, "EX", strconv.Itoa(*query.expiration))
	}

	args = append(args, query.cmd.Name)
	args = append(args, query.cmd.Args...)
	return newCmd("SETCHAN", args...)
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
