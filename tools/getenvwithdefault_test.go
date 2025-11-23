package tools

import (
	"os"
	"testing"
	"time"
)

func TestGetenvWithDefault_Time(t *testing.T) {
	defaultTime, _ := time.Parse("2006-01-02", "1900-01-01")
	checkTimeString := "2011-11-11"
	checkTime, _ := time.Parse("2006-01-02", checkTimeString)

	tests := []struct {
		name         string
		defaultValue time.Time
		envValue     string
		expected     time.Time
		format       []string
	}{
		{
			name:         "empty environment variable",
			defaultValue: defaultTime,
			envValue:     "",
			expected:     defaultTime,
			format:       []string{"2006-01-02"},
		},
		{
			name:         "valid environment variable",
			defaultValue: defaultTime,
			envValue:     checkTimeString,
			expected:     checkTime,
			format:       []string{"2006-01-02"},
		},
		{
			name:         "invalid environment variable",
			defaultValue: defaultTime,
			envValue:     "invalid",
			expected:     defaultTime,
			format:       []string{"2006-01-02"},
		},
		{
			name:         "invalid format",
			defaultValue: defaultTime,
			envValue:     checkTimeString,
			expected:     defaultTime,
			format:       []string{"invalid"},
		},
		{
			name:         "empty format",
			defaultValue: defaultTime,
			envValue:     checkTimeString,
			expected:     defaultTime,
			format:       []string{},
		},
		{
			name:         "multiple formats",
			defaultValue: defaultTime,
			envValue:     checkTimeString,
			expected:     checkTime,
			format:       []string{"invalid", "invalid", "2006-01-02"},
		},
	}

	/*
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				result := Unindent(tt.input)
				if result != tt.expected {
					t.Errorf("Unindent() = %q, expected %q", result, tt.expected)
				}
			})
		}
	*/
	for _, item := range tests {
		if item.envValue == "" {
			os.Unsetenv("TEST_TIME_VALUE")
		} else {
			os.Setenv("TEST_TIME_VALUE", item.envValue)
		}

		t.Run(item.name, func(t *testing.T) {
			result := GetenvWithDefault("TEST_TIME_VALUE", item.defaultValue, listOfAny(item.format)...)

			if result != item.expected {
				t.Errorf("given %s, expected %s, got %s", item.envValue, item.expected, result)
			}
		})
	}
}

// convert a slice of strings into a slice of any
func listOfAny(items []string) []any {
	result := make([]any, len(items))
	for i, v := range items {
		result[i] = v
	}
	return result
}
