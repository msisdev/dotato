package gardenpath

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const DefaultSeparator = '/'

var Root = GardenPath{""}

// GardenPath is a smart path representation.
// It is a sequence of directory names starting from root
// directory.
// 
// The first element represents the root directory:
//  - linux or else: empty string
//  - windows: volume name (e.g. C:)
type GardenPath []string

// New constructor handles:
//
//  1. Clean dot and double dot  
//  2. environment variable expansion
//  3. tilde replacement,
//  4. absolute path conversion,
//  5. trailing slash removal.
//
// It returns nil if the path is empty.
func New(path string) (GardenPath, error) {
	if path == "" {
		return nil, nil
	}

	// Clean dots ("." or "..")
	path = filepath.Clean(path)

	// Expand env vars
	path, notFound := expandEnv(path)
	if len(notFound) > 0 {
		return nil, fmt.Errorf("env vars not found: %v", notFound)
	}

	// Replace tilde
	path, err := expandTilde(path)
	if err != nil {
		return nil, err
	}

	// Resolve working directory
	path, err = filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// Remove trailing slash
	sep := string(os.PathSeparator)
	path = strings.TrimSuffix(path, sep)

	return strings.Split(path, sep), nil
}

// Get absolute path.
func (p GardenPath) Abs() string {
	return strings.Join(p, string(os.PathSeparator))
}

// Return the last element.
func (p GardenPath) Last() string {
	if len(p) == 0 {
		return ""
	}
	return p[len(p)-1]
}

// Return the parent path.
// Technically it returns [0:len(p)-1].
func (p GardenPath) Parent() GardenPath {
	if len(p) == 0 {
		return p
	}
	return p[:len(p)-1]
}

func (p GardenPath) IsEqual(other GardenPath) bool {
	if len(p) != len(other) {
		return false
	}
	for i := range p {
		if p[i] != other[i] {
			return false
		}
	}
	return true
}
