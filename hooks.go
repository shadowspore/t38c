package t38c

// DelHook remove a specified hook.
func (client *Client) DelHook(name string) error {
	return client.jExecute(nil, "DELHOOK", name)
}

// Hooks returns all hooks matching pattern.
func (client *Client) Hooks(pattern string) ([]Hook, error) {
	var resp struct {
		Hooks []Hook `json:"hooks"`
	}

	err := client.jExecute(&resp, "HOOKS", pattern)
	if err != nil {
		return nil, err
	}

	return resp.Hooks, nil
}

// PDelHook removes all hooks that match the specified pattern.
func (client *Client) PDelHook(pattern string) error {
	return client.jExecute(nil, "PDELHOOK", pattern)
}

// SetHook creates a webhook which points to a geofenced search.
// If a hook is already associated to that name, it’ll be overwritten.
func (client *Client) SetHook(name, endpoint string, query GeofenceQueryBuilder) SetHookQueryBuilder {
	return newSetHookQueryBuilder(client, name, endpoint, query.toCmd())
}
