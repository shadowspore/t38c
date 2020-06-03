package t38c

// JSetQueryBuilder struct
type JSetQueryBuilder struct {
	client   tile38Client
	key      string
	objectID string
	path     string
	value    string
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

func (query JSetQueryBuilder) toCmd() *tileCmd {
	cmd := newTileCmd("JSET", query.key, query.objectID, query.path, query.value)
	return cmd
}

// Do cmd
func (query JSetQueryBuilder) Do() error {
	cmd := query.toCmd()
	return query.client.jExecute(nil, cmd.Name, cmd.Args...)
}

// DoStr allows value to be interpreted as a string
func (query JSetQueryBuilder) DoStr() error {
	cmd := query.toCmd()
	cmd.appendArgs("STR")
	return query.client.jExecute(nil, cmd.Name, cmd.Args...)
}

// DoRaw allows value to be interpreted as a serialized JSON object
func (query JSetQueryBuilder) DoRaw() error {
	cmd := query.toCmd()
	cmd.appendArgs("RAW")
	return query.client.jExecute(nil, cmd.Name, cmd.Args...)
}
