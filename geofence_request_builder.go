package t38c

// GeofenceRequest struct
type GeofenceRequest struct {
	key        string
	target     string
	objectType string
	pattern    string
	meters     int
	actions    []string
	opts       []SearchOption
}

func NewFenceRequest(key, target, pattern string, meters int) *GeofenceRequest {
	return &GeofenceRequest{
		key:     key,
		target:  target,
		pattern: pattern,
		meters:  meters,
	}
}

func (req *GeofenceRequest) Actions(actions ...string) *GeofenceRequest {
	req.actions = actions
	return req
}

func (req *GeofenceRequest) WithOptions(opts ...SearchOption) *GeofenceRequest {
	req.opts = opts
	return req
}
