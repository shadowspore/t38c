package t38c

// JSetQueryBuilder struct
type JSetQueryBuilder struct {
	client   tile38Client
	key      string
	objectID string
	path     string
	value    string
	str      bool
	raw      bool
}

func newJSetQueryBuilder(client tile38Client, key, objectID, path, value string) JSetQueryBuilder {
	return JSetQueryBuilder{
		client:   client,
		key:      key,
		objectID: objectID,
		path:     path,
		value:    value,
	}
}

func (query JSetQueryBuilder) toCmd() cmd {
	args := []string{
		query.key, query.objectID, query.path, query.value,
	}

	if query.str {
		args = append(args, "STR")
	}

	if query.raw {
		args = append(args, "RAW")
	}
	return newCmd("JSET", args...)
}

// Do cmd
func (query JSetQueryBuilder) Do() error {
	cmd := query.toCmd()
	return query.client.jExecute(nil, cmd.Name, cmd.Args...)
}

// Str allows value to be interpreted as a string
func (query JSetQueryBuilder) Str() JSetQueryBuilder {
	query.str = true
	return query
}

// Raw allows value to be interpreted as a serialized JSON object
func (query JSetQueryBuilder) Raw() JSetQueryBuilder {
	query.raw = true
	return query
}
