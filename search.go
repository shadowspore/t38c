package t38c

import "encoding/json"

func (client *Tile38Client) Search(req *SearchRequest) (*SearchResponse, error) {
	resp := &SearchResponse{}
	cmd := req.BuildCommand()
	b, err := client.Execute(cmd.Name, cmd.Args...)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
