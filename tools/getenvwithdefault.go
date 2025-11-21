package tools

import (
	"os"
	"strconv"
	"time"
)

// GetenvWithDefault is a typed version of os.Getenv that allows you to
// specify a default value. The return type of the method is the type
// of the defaultValue argument. Currently supported types:
//
//   - string
//   - int
//   - bool
//   - float64
//   - time.Duration
//   - time.Time (must provide format as third argument)
//
// If you use any other type you will always get the default value.
func GetenvWithDefault[T any](name string, defaultValue T, options ...any) (value T) {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}

	var result any
	var zero T

	switch any(zero).(type) {
	case string:
		result = val
	case int:
		v, err := strconv.Atoi(val)
		if err != nil {
			return defaultValue
		}
		result = v
	case bool:
		v, err := strconv.ParseBool(val)
		if err != nil {
			return defaultValue
		}
		result = v
	case float64:
		v, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return defaultValue
		}
		result = v
	case time.Duration:
		v, err := time.ParseDuration(val)
		if err != nil {
			return defaultValue
		}
		result = v
	case time.Time:
		if len(options) == 0 {
			return defaultValue
		}
		format, ok := options[0].(string)
		if !ok {
			return defaultValue
		}
		v, err := time.Parse(format, val)
		if err != nil {
			return defaultValue
		}
		result = v
	default:
		return defaultValue
	}

	return result.(T)
}
