package gardenpath

import (
	"os"
	"os/user"
	"strings"
)

// A modification from https://pkg.go.dev/github.com/go-git/go-git/v5/plumbing/format/gitignore
func expandTilde(rawPath string) (string, error) {
	if !strings.HasPrefix(rawPath, "~") {
		return rawPath, nil
	}

	// Is tilde followed by a slash?
	end := strings.Index(rawPath, "/")
	if end == -1 {
		end = len(rawPath)	// No slash: use the whole string
	}

	// Path in ~/
	if end == 1 {
		home, err := os.UserHomeDir()
		if err != nil {
			return rawPath, err
		}
		return strings.Replace(rawPath, "~", home, 1), nil
	}

	// Path in ~<username>
	username := rawPath[1:end]
	u, err := user.Lookup(username)
	if err != nil {
		return rawPath, err
	}
	return strings.Replace(rawPath, rawPath[:end], u.HomeDir, 1), nil
}

// expandEnv returns the expanded string and
// a slice of environment variables that were not found.
func expandEnv(s string) (string, []string) {
	// Not found env vars
	notFound := make([]string, 0)

	// Expand env vars
	expanded := os.Expand(s, func(env string) string {
		// Does env var exist?
		val, ok := os.LookupEnv(env)
		if !ok {
			notFound = append(notFound, env)
		}
		return val
	})

	// 
	if len(notFound) > 0 {
		return "", notFound
	}
	return expanded, nil
}
