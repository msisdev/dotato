package gardenpath

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var ErrEnvVarNotSet = fmt.Errorf("env var not set")

// GardenPath is a smart path representation.
// It is a sequence of directory names starting from root
// directory.
//
// The first element represents the root directory:
//   - linux or else: empty string
//   - windows: volume name (e.g. C:)
type GardenPath []string

// New constructor handles:
//
//  1. Clean dot and double dot
//  2. tilde replacement,
//  3. absolute path conversion,
//  4. trailing slash removal.
//
// It returns nil if the path is empty.
func New(path string) (GardenPath, error) {
	gp, _, err := NewCheckEnv(path)
	return gp, err
}

// Returns a list of env vars that were not found.
func NewCheckEnv(path string) (gp GardenPath, notFound []string, err error) {
	if path == "" {
		return nil, nil, nil
	}

	// Clean dots ("." or "..")
	path = filepath.Clean(path)

	// (linux) Replace tilde
	path, err = expandTilde(path)
	if err != nil {
		return
	}

	// Resolve working directory
	path, err = filepath.Abs(path)
	if err != nil {
		return
	}

	// Remove trailing slash
	sep := string(os.PathSeparator)
	path = strings.TrimSuffix(path, sep)

	// (windows) handle volume name
	if vol := filepath.VolumeName(path); vol != "" {
		path = strings.TrimPrefix(path, vol) // remove volume name
		path = strings.TrimPrefix(path, sep) // remove leading separator
		gp = GardenPath{vol}
		if path != "" {
			gp = append(gp, strings.Split(path, sep)...)
		}
		return
	}

	gp = strings.Split(path, sep)
	return
}

// Get absolute path.
func (p GardenPath) Abs() string {
	// handle empty path
	if len(p) == 0 {
		return ""
	}

	if len(p) == 1 {
		if p[0] == "" {
			// (linux) handle root directory
			return "/"
		} else {
			// (windows) handle volume name
			return p[0] + string(os.PathSeparator)
		}
	}

	return strings.Join(p, string(os.PathSeparator))
}

func (p GardenPath) Copy() GardenPath {
	cp := make(GardenPath, len(p))
	copy(cp, p)
	return cp
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
