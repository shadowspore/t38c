package t38c

import (
	"strconv"
	"strings"
)

// HookBuilder struct
type HookBuilder struct {
	Name      string
	Endpoints []string
	Metas     []Meta
	Command   Command
	Ex        *int
}

// Args ...
func (hook *HookBuilder) Args() []string {
	var args []string
	args = append(args, hook.Name)
	args = append(args, strings.Join(hook.Endpoints, ","))

	for _, meta := range hook.Metas {
		args = append(args, "META")
		args = append(args, meta.Name)
		args = append(args, meta.Value)
	}

	if hook.Ex != nil {
		args = append(args, "EX")
		args = append(args, strconv.Itoa(*hook.Ex))
	}

	args = append(args, hook.Command.Name)
	args = append(args, hook.Command.Args...)
	return args
}

// NewHook ...
func NewHook(name string, endpoints []string, req GeofenceRequestable) *HookBuilder {
	return &HookBuilder{
		Name:      name,
		Endpoints: endpoints,
		Command:   req.GeofenceCommand(),
	}
}

// Meta ...
func (hook *HookBuilder) Meta(name, value string) *HookBuilder {
	hook.Metas = append(hook.Metas, Meta{
		Name:  name,
		Value: value,
	})

	return hook
}

// Expiration ...
func (hook *HookBuilder) Expiration(seconds int) *HookBuilder {
	hook.Ex = &seconds
	return hook
}

// DelHook remove a specified hook.
func (client *Client) DelHook(name string) error {
	return client.JExecute(nil, "DELHOOK", name)
}

// Hooks returns all hooks matching pattern.
func (client *Client) Hooks(pattern string) ([]Hook, error) {
	var resp struct {
		Hooks []Hook `json:"hooks"`
	}

	err := client.JExecute(&resp, "HOOKS", pattern)
	if err != nil {
		return nil, err
	}

	return resp.Hooks, nil
}

// PDelHook removes all hooks that match the specified pattern.
func (client *Client) PDelHook(pattern string) error {
	return client.JExecute(nil, "PDELHOOK", pattern)
}

// SetHook creates a webhook which points to a geofenced search.
// If a hook is already associated to that name, it’ll be overwritten.
func (client *Client) SetHook(hook *HookBuilder) error {
	return client.JExecute(nil, "SETHOOK", hook.Args()...)
}
