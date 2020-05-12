package t38c

import "strconv"

// SetChannelQueryBuilder struct
type SetChannelQueryBuilder struct {
	client     *Client
	name       string
	cmd        Command
	metas      []Meta
	expiration *int
}

func newSetChannelQueryBuilder(client *Client, name string, cmd Command) SetChannelQueryBuilder {
	return SetChannelQueryBuilder{
		client: client,
		name:   name,
		cmd:    cmd,
	}
}

func (query SetChannelQueryBuilder) toCmd() Command {
	args := []string{query.name}
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
	return NewCommand("SETCHAN", args...)
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
