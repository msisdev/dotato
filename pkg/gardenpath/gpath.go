package gardenpath

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Path represents an absolute path as a sequence of directories.
// It always starts with an empty string to distinguish nil from
// root directory.
// Therefore the length is always bigger than 0.
type GardenPath []string

// NewGardenPath handles:
// (1) environment variable expansion
// (2) tilde replacement,
// (3) absolute path conversion,
// (4) trailing slash removal.
func New(path string) (GardenPath, error) {
	// Expand env vars
	path, notFound := expandEnv(path)
	if len(notFound) > 0 {
		return nil, fmt.Errorf("env vars not found: %v", notFound)
	}

	// Replace tilde
	path, err := replaceTilde(path)
	if err != nil {
		return nil, err
	}

	// Insert PWD
	if !strings.HasPrefix(path, "/") {
		path = filepath.Join(os.Getenv("PWD"), path)
	}

	// Remove trailing slash
	path = strings.TrimSuffix(path, "/")

	return strings.Split(path, "/"), nil
}

// Return absolute path
func (p GardenPath) String() string {
	return strings.Join(p, "/")
}
