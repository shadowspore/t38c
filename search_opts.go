package t38c

import "strconv"

type whereOpt struct {
	Field string
	Min   float64
	Max   float64
}

type whereinOpt struct {
	Field  string
	Values []float64
}

type whereEvalOpt struct {
	Name  string
	IsSHA bool
	Args  []string
}

type searchOpts struct {
	Asc       bool
	Desc      bool
	NoFields  bool
	Clip      bool
	Distance  bool
	Cursor    *int
	Limit     *int
	Sparse    *int
	Where     []whereOpt
	Wherein   []whereinOpt
	Match     []string
	WhereEval []whereEvalOpt
	RawQuery  []string
}

func (opts searchOpts) Args() (args []string) {
	for _, opt := range opts.Where {
		args = append(args, "WHERE", opt.Field, floatString(opt.Min), floatString(opt.Max))
	}

	for _, opt := range opts.RawQuery {
		args = append(args, "WHERE", opt)
	}

	for _, opt := range opts.Wherein {
		values := make([]string, len(opt.Values))
		for i, val := range opt.Values {
			values[i] = floatString(val)
		}
		args = append(args, "WHEREIN", opt.Field)
		args = append(args, strconv.Itoa(len(opt.Values)))
		args = append(args, values...)
	}

	for _, opt := range opts.WhereEval {
		if opt.IsSHA {
			args = append(args, "WHEREEVALSHA", opt.Name)
		} else {
			args = append(args, "WHEREEVAL", opt.Name)
		}
		args = append(args, strconv.Itoa(len(opt.Args)))
		args = append(args, opt.Args...)
	}

	for _, pattern := range opts.Match {
		args = append(args, "MATCH", pattern)
	}

	if opts.Asc {
		args = append(args, "ASC")
	}

	if opts.Desc {
		args = append(args, "DESC")
	}

	if opts.NoFields {
		args = append(args, "NOFIELDS")
	}

	if opts.Clip {
		args = append(args, "CLIP")
	}

	if opts.Distance {
		args = append(args, "DISTANCE")
	}

	if opts.Cursor != nil {
		args = append(args, "CURSOR", strconv.Itoa(*opts.Cursor))
	}

	if opts.Limit != nil {
		args = append(args, "LIMIT", strconv.Itoa(*opts.Limit))
	}

	if opts.Sparse != nil {
		args = append(args, "SPARSE", strconv.Itoa(*opts.Sparse))
	}
	return
}
