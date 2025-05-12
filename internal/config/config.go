package config

import (
	"fmt"

	"github.com/goccy/go-yaml"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

var (
	ErrVersionNotFound  = fmt.Errorf("version not found")
	ErrModeNotFound     = fmt.Errorf("mode not found")
	ErrGroupNotFound    = fmt.Errorf("group not found")
	ErrResolverNotFound = fmt.Errorf("resolver not found")
)

const (
	ModeFile    = "file"
	ModeLink    = "link"
	ModeDefault = ModeFile

	Version1      = "v1"
	ConfigVersion = Version1
)

type Config struct {
	Version string                       `yaml:"version"`
	Mode    string                       `yaml:"mode"`
	Plans   map[string][]string          `yaml:"plans"`
	Groups  map[string]map[string]string `yaml:"groups"`
}

func New() *Config {
	return &Config{
		Version: ConfigVersion,
		Mode:    ModeDefault,
		Plans:   map[string][]string{},
		Groups:  map[string]map[string]string{},
	}
}

func NewFromByte(data []byte) (*Config, error) {
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	// Check required fields
	if config.Version == "" {
		return nil, ErrVersionNotFound
	}
	if config.Mode == "" {
		return nil, ErrModeNotFound
	}

	return &config, nil
}

func NewFromString(str string) (*Config, error) {
	return NewFromByte([]byte(str))
}

// For testing purpose
func (r Config) IsEqual(other *Config) bool {
	// Compare versions
	if r.Version != other.Version {
		return false
	}

	// Compare modes
	if r.Mode != other.Mode {
		return false
	}

	// Compare plans
	if len(r.Plans) != len(other.Plans) {
		return false
	}
	for key, plan := range r.Plans {
		if !cmpStrSlice(plan, other.Plans[key]) {
			return false
		}
	}

	// Compare groups
	if len(r.Groups) != len(other.Groups) {
		return false
	}
	for group, resolvers := range r.Groups {
		if len(resolvers) != len(other.Groups[group]) {
			return false
		}
		for key, resolver := range resolvers {
			if other.Groups[group][key] != resolver {
				return false
			}
		}
	}

	return true
}

// Returns groups of the plan as a map.
//
// If a plan has empty list, it returns all groups.
func (c Config) GetGroups(plan string) map[string]bool {
	// Find list by plan
	list, ok := c.Plans[plan]
	if !ok {
		return nil
	}

	set := make(map[string]bool)

	if list == nil {
		// Empty list means all
		for group := range c.Groups {
			set[group] = true
		}
	} else {
		// Convert list to set
		for _, group := range list {
			set[group] = true
		}
	}

	return set
}

func (c Config) GetGroupBase(
	group, resolver string,
) (
	base gp.GardenPath,
	notFound []string,
	err error,
) {
	// Get group resolvers with group name
	resolverMap, ok := c.Groups[group]
	if !ok {
		err = ErrGroupNotFound
		return
	}

	// Get raw path with resolver
	rawPath, ok := resolverMap[resolver]
	if !ok {
		err = ErrResolverNotFound
		return
	}

	// Expand env vars
	rawPath, notFound = expandEnv(rawPath)
	if len(notFound) > 0 {
		err = fmt.Errorf("env var is not set in %s", rawPath)
		return
	}

	base, notFound, err = gp.NewCheckEnv(rawPath)
	return
}

// Returns:
//   - groups - map of group and base pairs
//   - notFound - env vars that are not found in the base
//   - err - error if any
func (c Config) GetGroupBaseAll(resolver string) (groups map[string]gp.GardenPath, notFound []string, err error) {
	groups = make(map[string]gp.GardenPath)

	// Make groups
	for group, resolvers := range c.Groups {
		// Find resolver
		path, ok := resolvers[resolver]
		if !ok {
			continue // Resolver not found
		}

		// Get garden path
		var (
			gpath gp.GardenPath
			envs  []string
		)
		gpath, envs, err = gp.NewCheckEnv(path)
		if err != nil {
			if err == gp.ErrEnvVarNotSet {
				notFound = append(notFound, envs...) // Append not found env vars
			} else {
				return
			}
		}

		// Add to groups
		groups[group] = gpath
	}

	return
}
