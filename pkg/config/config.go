package config

import (
	"fmt"

	"github.com/goccy/go-yaml"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

var (
	ErrVersionNotFound 	= fmt.Errorf("version not found")
	ErrModeNotFound    	= fmt.Errorf("mode not found")
)

type Mode string
const (
	ModeFile Mode = "file"
	ModeLink Mode = "link"
	ModeDefault Mode = ModeFile
)

type Config struct {
	Version string               					`yaml:"version"`
	Mode		Mode               						`yaml:"mode"`
	Plans   map[string][]string						`yaml:"plans"`
	Groups  map[string]map[string]string	`yaml:"groups"`
}

func New() *Config {
	return &Config{
		Version:	DotatoVersion(),
		Mode:			ModeDefault,
		Plans:   	map[string][]string{},
		Groups:  	map[string]map[string]string{},
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

// Returns groups of the plan as a map
func (c Config) GetPlan(plan string) map[string]bool {
	// Does the plan exist?
	if _, ok := c.Plans[plan]; !ok {
		return nil
	}

	// Convert plan to a map
	groups := make(map[string]bool)
	for _, group := range c.Plans[plan] {
		groups[group] = true
	}

	return groups
}

// Returns:
//  - groups - groups that has the resolver, with garden path
//  - notFound - env vars that are not found in the resolver
//	- err - error if any
func (c Config) GetGroups(resolver string) (groups map[string]gp.GardenPath, notFound []string, err error) {
	groups = make(map[string]gp.GardenPath)
	
	// Make groups
	for group, resolvers := range c.Groups {
		// Find resolver
		path, ok := resolvers[resolver]
		if !ok {
			continue	// Resolver not found
		}

		// Get garden path
		var (
			gpath gp.GardenPath
			envs []string
		)
		gpath, envs, err = gp.NewCheckEnv(path)
		if err != nil {
			if err == gp.ErrEnvVarNotFound {
				notFound = append(notFound, envs...)	// Append not found env vars
			} else {
				return
			}
		}

		// Add to groups
		groups[group] = gpath
	}

	return
}
