package gardenpath

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	DefaultSeparator = '/'
)

// GardenPath is a smart path representation.
// It is a sequence of directory names starting from root
// directory.
// 
// It always contains an empty string at index 0
// to distinguish root directory from nil.
type GardenPath []string

// NewGardenPath handles:
//
//  1. Clean dot and double dot  
//  2. environment variable expansion
//  3. tilde replacement,
//  4. absolute path conversion,
//  5. trailing slash removal.
//
// Returns nil if the path is empty.
func New(path string) (GardenPath, error) {
	return NewWithSep(path, DefaultSeparator)
}

func NewWithSep(path string, separator rune) (GardenPath, error) {
	if path == "" {
		return nil, nil
	}

	sep := string(separator)

	// Clean dots ("." or "..")
	path = filepath.Clean(path)

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
	if !strings.HasPrefix(path, sep) {
		path = filepath.Join(os.Getenv("PWD"), path)
	}

	// Remove trailing slash
	path = strings.TrimSuffix(path, sep)

	return strings.Split(path, sep), nil
}

// Return absolute path.
func (p GardenPath) String() string {
	return strings.Join(p, string(DefaultSeparator))
}

// Return the last element.
func (p GardenPath) Last() string {
	if len(p) == 0 {
		return ""
	}
	return p[len(p)-1]
}

// Return the parent path.
// This works with both directory and file paths.
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
