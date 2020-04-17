package t38c

// SearchRequest struct
type SearchRequest struct {
	Cmd           string
	Key           string
	OutputFormat  OutputFormat
	Area          Command
	SearchOptions []SearchOption
}

// BuildCommand builds tile38 search command.
func (req *SearchRequest) BuildCommand() Command {
	var args []string
	args = append(args, req.Key)

	for _, opt := range req.SearchOptions {
		args = append(args, opt.Name)
		args = append(args, opt.Args...)
	}

	if len(req.OutputFormat.Name) > 0 {
		args = append(args, req.OutputFormat.Name)
		args = append(args, req.OutputFormat.Args...)
	}

	args = append(args, req.Area.Name)
	args = append(args, req.Area.Args...)

	return NewCommand(req.Cmd, args...)
}

// Format set response format.
func (req *SearchRequest) Format(fmt OutputFormat) *SearchRequest {
	req.OutputFormat = fmt
	return req
}

// WithOptions sets the optional parameters for request.
func (req *SearchRequest) WithOptions(opts ...SearchOption) *SearchRequest {
	req.SearchOptions = opts
	return req
}

// Within searches a collection for objects that are fully contained inside of a specified bounding area.
func Within(key string, area SearchArea) *SearchRequest {
	return &SearchRequest{
		Cmd:  "WITHIN",
		Key:  key,
		Area: Command(area),
	}
}

// Intersects searches a collection for objects that intersect a specified bounding area.
func Intersects(key string, area SearchArea) *SearchRequest {
	return &SearchRequest{
		Cmd:  "INTERSECTS",
		Key:  key,
		Area: Command(area),
	}
}

// Nearby command searches a collection for objects that are close to a specified point.
// The KNN algorithm is used instead of the standard overlap+Haversine algorithm,
// sorting the results in order of ascending distance from that point, i.e., nearest first.
func Nearby(key string, lat, lon, meters float64) *SearchRequest {
	return &SearchRequest{
		Cmd:  "NEARBY",
		Key:  key,
		Area: NewCommand("POINT", floatString(lat), floatString(lon), floatString(meters)),
	}
}

// Search iterates though a key’s string values.
func Search(key string) *SearchRequest {
	return &SearchRequest{
		Cmd: "SEARCH",
		Key: key,
	}
}

// Scan incrementally iterates though a key.
func Scan(key string) *SearchRequest {
	return &SearchRequest{
		Cmd: "SCAN",
		Key: key,
	}
}

// Search execute a search request.
func (client *Client) Search(req *SearchRequest) (*SearchResponse, error) {
	cmd := req.BuildCommand()

	resp := &SearchResponse{}
	err := client.jExecute(&resp, cmd.Name, cmd.Args...)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
