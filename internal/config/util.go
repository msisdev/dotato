package config

import (
	"os"
	"sort"
)

func cmpStrSlice(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Returns the expanded string and a list of
// missing env vars.
func expandEnv(s string) (expanded string, notFound []string) {
	expanded = os.Expand(s, func(env string) string {
		val, ok := os.LookupEnv(env)
		if !ok {
			notFound = append(notFound, env)
		}

		return val
	})

	return
}
