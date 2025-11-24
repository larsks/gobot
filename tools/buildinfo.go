package tools

import (
	"runtime/debug"
)

// BuildInfoMap converts BuildInfo.Settings into a map[string]string.
func BuildInfoMap(bi *debug.BuildInfo) map[string]string {
	res := make(map[string]string)
	for _, setting := range bi.Settings {
		res[setting.Key] = setting.Value
	}
	return res
}
