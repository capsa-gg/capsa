package util

import (
	"fmt"
	"time"
)

// ExtractTimeFromAny tries to extract a time.Time from an any value.
func ExtractTimeFromAny(in any) (time.Time, error) {
	switch v := in.(type) {
	case time.Time:
		return v, nil
	case *time.Time:
		return *v, nil
	case string:
		// Assuming the string is in a standard time format
		parsedTime, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return time.Time{}, err
		}

		return parsedTime, nil
	case int64:
		// Assuming the int64 is a Unix timestamp in seconds
		return time.Unix(v, 0), nil
	case float64:
		// Assuming the float64 is a Unix timestamp in seconds
		return time.Unix(int64(v), 0), nil
	default:
		return time.Time{}, fmt.Errorf("unsupported type: %T", v)
	}
}
