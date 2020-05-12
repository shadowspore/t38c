package t38c

import (
	"encoding/json"
	"fmt"
	"time"
)

// StringTime struct
type StringTime struct {
	time.Time
}

// UnmarshalJSON ...
func (st *StringTime) UnmarshalJSON(data []byte) error {
	runes := []rune(string(data))
	if len(runes) < 3 {
		return fmt.Errorf("bad time response: %s", data)
	}

	// remove quotes
	str := string(runes[1 : len(runes)-1])

	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return err
	}

	st.Time = t
	return nil
}

// MarshalJSON ...
func (st *StringTime) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(st.Time)
	if err != nil {
		return nil, err
	}

	return []byte("\"" + string(b) + "\""), nil
}
