package t38c

import "strconv"

// SearchOption ...
type SearchOption Command

var (
	// Asc ...
	Asc = SearchOption(NewCommand("ASC"))

	// Desc ...
	Desc = SearchOption(NewCommand("DESC"))

	// Count ...
	Count = func(count int) SearchOption {
		return SearchOption(NewCommand("COUNT", strconv.Itoa(count)))
	}

	// NoFields ...
	NoFields = SearchOption(NewCommand("NOFIELDS"))

	// Clip ...
	Clip = SearchOption(NewCommand("CLIP"))

	// Cursor ...
	Cursor = func(start int) SearchOption {
		return SearchOption(NewCommand("CURSOR", strconv.Itoa(start)))
	}

	// Limit ...
	Limit = func(count int) SearchOption {
		return SearchOption(NewCommand("LIMIT", strconv.Itoa(count)))
	}

	// Sparse will distribute the results of a search evenly across the requested area.
	Sparse = func(n int) SearchOption {
		return SearchOption(NewCommand("SPARSE", strconv.Itoa(n)))
	}

	// Where allows for filtering out results based on field values.
	Where = func(field string, min, max float64) SearchOption {
		return SearchOption(NewCommand("WHERE", field, floatString(min), floatString(max)))
	}

	// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
	Wherein = func(field string, values ...float64) SearchOption {
		var args []string
		args = append(args, strconv.Itoa(len(values)))
		for _, val := range values {
			args = append(args, floatString(val))
		}

		return SearchOption(NewCommand("WHEREIN", args...))
	}

	// Match is similar to WHERE except that it works on the object id instead of fields.
	Match = func(pattern string) SearchOption {
		return SearchOption(NewCommand("MATCH", pattern))
	}

	// Distance ...
	Distance = SearchOption(NewCommand("DISTANCE"))
)

// SetOption ...
type SetOption Command

var (
	// Field ...
	SetField = func(name string, value float64) SetOption {
		return SetOption(NewCommand("FIELD", name, floatString(value)))
	}

	// Expiration ...
	Expiration = func(seconds int) SetOption {
		return SetOption(NewCommand("EX", strconv.Itoa(seconds)))
	}

	// IfNotExists ...
	IfNotExists = SetOption(NewCommand("NX"))

	// IfExists ...
	IfExists = SetOption(NewCommand("XX"))
)
