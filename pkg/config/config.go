package config

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

type Mode string

const (
	ModeFile Mode = "file"
	ModeLink Mode = "link"
	ModeDefault = ModeFile
)

var (
	ErrVersionNotFound = fmt.Errorf("version not found")
	ErrModeNotFound    = fmt.Errorf("mode not found")
)

// YAML schema
type Config struct {
	Version string               	`yaml:"version"`
	Mode		Mode               		`yaml:"mode"`
	Plans   map[string]GroupList	`yaml:"plans"`
	Groups  map[string]string    	`yaml:"groups"`
}

func NewConfig() *Config {
	return &Config{
		Version:	GetDotatoVersion(),
		Mode:			ModeDefault,
		Plans:   	map[string]GroupList{},
		Groups:  	map[string]string{}, // sample group
	}
}

func NewConfigFromByte(data []byte) (*Config, error) {
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

func NewConfigFromStr(str string) (*Config, error) {
	return NewConfigFromByte([]byte(str))
}

// IsEqual can compare two Config objects.
// It is for testing purpose.
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
	for key, rawPlan := range r.Plans {
		if !rawPlan.IsEqual(other.Plans[key]) {
			return false
		}
	}

	// Compare groups
	if len(r.Groups) != len(other.Groups) {
		return false
	}
	for key, rawBase := range r.Groups {
		if rawBase != other.Groups[key] {
			return false
		}
	}

	return true
}
