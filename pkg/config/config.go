package config

import (
	"fmt"

	"github.com/goccy/go-yaml"
	// gp "github.com/msisdev/dotato/pkg/gardenpath"
)

var (
	ErrVersionNotFound = fmt.Errorf("version not found")
	ErrModeNotFound    = fmt.Errorf("mode not found")
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
