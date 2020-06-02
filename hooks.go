package t38c

// Hooks struct
type Hooks struct {
	client *Client
}

// DelHook remove a specified hook.
func (hooks *Hooks) DelHook(name string) error {
	return hooks.client.jExecute(nil, "DELHOOK", name)
}

// Hooks returns all hooks matching pattern.
func (hooks *Hooks) Hooks(pattern string) ([]Hook, error) {
	var resp struct {
		Hooks []Hook `json:"hooks"`
	}

	err := hooks.client.jExecute(&resp, "HOOKS", pattern)
	if err != nil {
		return nil, err
	}

	return resp.Hooks, nil
}

// PDelHook removes all hooks that match the specified pattern.
func (hooks *Hooks) PDelHook(pattern string) error {
	return hooks.client.jExecute(nil, "PDELHOOK", pattern)
}

// SetHook creates a webhook which points to a geofenced search.
// If a hook is already associated to that name, it’ll be overwritten.
func (hooks *Hooks) SetHook(name, endpoint string, query GeofenceQueryBuilder) SetHookQueryBuilder {
	return newSetHookQueryBuilder(hooks.client, name, endpoint, query.toCmd())
}
