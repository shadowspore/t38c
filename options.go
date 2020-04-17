package t38c

import "strconv"

// SearchOption ...
type SearchOption Command

var (
	// Asc order. Only for SEARCH and SCAN commands.
	Asc = SearchOption(NewCommand("ASC"))

	// Desc order. Only for SEARCH and SCAN commands.
	Desc = SearchOption(NewCommand("DESC"))

	// NoFields tells the server that you do not want field values returned with the search results.
	NoFields = SearchOption(NewCommand("NOFIELDS"))

	// Clip tells the server to clip intersecting objects by the bounding box area of the search.
	// It can only be used with these area formats: BOUNDS, TILE, QUADKEY, HASH.
	Clip = SearchOption(NewCommand("CLIP"))

	// Cursor is used to iterate though many objects from the search results.
	// An iteration begins when the CURSOR is set to Zero or not included with the request,
	// and completes when the cursor returned by the server is Zero.
	Cursor = func(start int) SearchOption {
		return SearchOption(NewCommand("CURSOR", strconv.Itoa(start)))
	}

	// Limit can be used to limit the number of objects returned for a single search request.
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
	// There can be multiple MATCH options in a single search.
	// The MATCH value is a simple glob pattern.
	Match = func(pattern string) SearchOption {
		return SearchOption(NewCommand("MATCH", pattern))
	}

	// Distance allows to return between objects. Only for NEARBY command.
	Distance = SearchOption(NewCommand("DISTANCE"))
)

// SetOption ...
type SetOption Command

var (
	// Field are extra data which belongs to an object.
	// A field is always a double precision floating point.
	// There is no limit to the number of fields that an object can have.
	Field = func(name string, value float64) SetOption {
		return SetOption(NewCommand("FIELD", name, floatString(value)))
	}

	// Expiration set the specified expire time, in seconds.
	Expiration = func(seconds int) SetOption {
		return SetOption(NewCommand("EX", strconv.Itoa(seconds)))
	}

	// IfNotExists only set the object if it does not already exist.
	IfNotExists = SetOption(NewCommand("NX"))

	// IfExists only set the object if it already exist.
	IfExists = SetOption(NewCommand("XX"))
)
