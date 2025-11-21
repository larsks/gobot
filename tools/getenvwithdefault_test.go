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
		defaultValue time.Time
		envValue     string
		expected     time.Time
		format       []string
	}{
		{
			defaultValue: defaultTime,
			envValue:     "",
			expected:     defaultTime,
			format:       []string{"2006-01-02"},
		},
		{
			defaultValue: defaultTime,
			envValue:     checkTimeString,
			expected:     checkTime,
			format:       []string{"2006-01-02"},
		},
		{
			defaultValue: defaultTime,
			envValue:     "invalid",
			expected:     defaultTime,
			format:       []string{"2006-01-02"},
		},
		{
			defaultValue: defaultTime,
			envValue:     checkTimeString,
			expected:     defaultTime,
			format:       []string{"invalid"},
		},
		{
			defaultValue: defaultTime,
			envValue:     checkTimeString,
			expected:     defaultTime,
			format:       []string{},
		},
		{
			defaultValue: defaultTime,
			envValue:     checkTimeString,
			expected:     checkTime,
			format:       []string{"invalid", "invalid", "2006-01-02"},
		},
	}

	for _, item := range tests {
		if item.envValue == "" {
			os.Unsetenv("TEST_TIME_VALUE")
		} else {
			os.Setenv("TEST_TIME_VALUE", item.envValue)
		}

		result := GetenvWithDefault("TEST_TIME_VALUE", item.defaultValue, listOfAny(item.format)...)

		if result != item.expected {
			t.Errorf("given %s, expected %s, got %s", item.envValue, item.expected, result)
		}
	}
}

func listOfAny(items []string) []any {
	result := make([]any, len(items))
	for i, v := range items {
		result[i] = v
	}
	return result
}
