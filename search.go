package t38c

import "encoding/json"

// Search ...
func (client *Tile38Client) Search(req *SearchRequest) (*SearchResponse, error) {
	cmd := req.BuildCommand()
	b, err := client.Execute(cmd.Name, cmd.Args...)
	if err != nil {
		return nil, err
	}

	resp := &SearchResponse{}
	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
