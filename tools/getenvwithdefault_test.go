package tools

import (
	"os"
	"testing"
	"time"
)

func TestGetenvWithDefault_Time(t *testing.T) {
	defaultTime, _ := time.Parse("2006-01-02", "1900-01-01")
	tests := []struct {
		defaultValue time.Time
		envValue     string
		expected     time.Time
		format       string
	}{
		{
			defaultValue: defaultTime,
			envValue:     "",
			expected:     defaultTime,
			format:       "2006-01-02",
		},
		{
			defaultValue: defaultTime,
			envValue:     "2011-11-11",
			expected: func() time.Time {
				t, _ := time.Parse("2006-01-02", "2011-11-11")
				return t
			}(),
			format: "2006-01-02",
		},
		{
			defaultValue: defaultTime,
			envValue:     "invalid",
			expected:     defaultTime,
			format:       "2006-01-02",
		},
		{
			defaultValue: defaultTime,
			envValue:     "2011-11-11",
			expected:     defaultTime,
			format:       "invalid",
		},
		{
			defaultValue: defaultTime,
			envValue:     "2011-11-11",
			expected:     defaultTime,
			format:       "",
		},
	}

	for _, item := range tests {
		if item.envValue == "" {
			os.Unsetenv("TEST_TIME_VALUE")
		} else {
			os.Setenv("TEST_TIME_VALUE", item.envValue)
		}

		var result time.Time
		if item.format == "" {
			result = GetenvWithDefault("TEST_TIME_VALUE", item.defaultValue)
		} else {
			result = GetenvWithDefault("TEST_TIME_VALUE", item.defaultValue, item.format)
		}

		if result != item.expected {
			t.Errorf("given %s, expected %s, got %s", item.envValue, item.expected, result)
		}
	}
}
