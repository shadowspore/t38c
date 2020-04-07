package t38c

import "strconv"

// SearchOption ...
type SearchOption Command

// Desc ...
func Desc() SearchOption {
	return SearchOption(NewCommand("DESC"))
}

// Asc ...
func Asc() SearchOption {
	return SearchOption(NewCommand("ASC"))
}

// Count ...
func Count(count int) SearchOption {
	return SearchOption(NewCommand("COUNT", strconv.Itoa(count)))
}

// NoFields ...
func NoFields() SearchOption {
	return SearchOption(NewCommand("NOFIELDS"))
}

// Clip ...
func Clip() SearchOption {
	return SearchOption(NewCommand("CLIP"))
}

// Cursor ...
func Cursor(start int) SearchOption {
	return SearchOption(NewCommand("CURSOR", strconv.Itoa(start)))
}

// Limit ...
func Limit(count int) SearchOption {
	return SearchOption(NewCommand("LIMIT", strconv.Itoa(count)))
}

// Sparse will distribute the results of a search evenly across the requested area.
func Sparse(n int) SearchOption {
	return SearchOption(NewCommand("SPARSE", strconv.Itoa(n)))
}

// Where allows for filtering out results based on field values.
func Where(field string, min, max float64) SearchOption {
	return SearchOption(NewCommand("WHERE", field, floatString(min), floatString(max)))
}

// Wherein is similar to Where except that it checks whether the object’s field value is in a given list.
func Wherein(field string, values ...float64) SearchOption {
	var args []string
	args = append(args, strconv.Itoa(len(values)))
	for _, val := range values {
		args = append(args, floatString(val))
	}

	return SearchOption(NewCommand("WHEREIN", args...))
}

// Match is similar to WHERE except that it works on the object id instead of fields.
func Match(pattern string) SearchOption {
	return SearchOption(NewCommand("MATCH", pattern))
}

// Distance ...
func Distance() SearchOption {
	return SearchOption(NewCommand("DISTANCE"))
}

// SetOption ...
type SetOption Command

// SetField ...
func SetField(name string, value float64) SetOption {
	return SetOption(NewCommand("FIELD", name, floatString(value)))
}

// func SetEX(d time.Duration) SetOption {
// 	return SetOption(
// 		"EX ",
// 	)
// }

// SetNX ...
func SetNX() SetOption {
	return SetOption(NewCommand("NX"))
}

// SetXX ...
func SetXX() SetOption {
	return SetOption(NewCommand("XX"))
}
