package t38c

type SearchRequest struct {
	Cmd           string
	Key           string
	OutputFormat  OutputFormat
	Area          Command
	SearchOptions []SearchOption
}

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

func (req *SearchRequest) Format(fmt OutputFormat) *SearchRequest {
	req.OutputFormat = fmt
	return req
}

func (req *SearchRequest) WithOptions(opts ...SearchOption) *SearchRequest {
	req.SearchOptions = opts
	return req
}

func Within(key string, area SearchArea) *SearchRequest {
	return &SearchRequest{
		Cmd:  "WITHIN",
		Key:  key,
		Area: Command(area),
	}
}

func Intersects(key string, area SearchArea) *SearchRequest {
	return &SearchRequest{
		Cmd:  "INTERSECTS",
		Key:  key,
		Area: Command(area),
	}
}

func Nearby(key string, area NearbyArea) *SearchRequest {
	return &SearchRequest{
		Cmd:  "NEARBY",
		Key:  key,
		Area: Command(area),
	}
}

func Search(key string) *SearchRequest {
	return &SearchRequest{
		Cmd: "SEARCH",
		Key: key,
	}
}

func Scan(key string) *SearchRequest {
	return &SearchRequest{
		Cmd: "SCAN",
		Key: key,
	}
}
