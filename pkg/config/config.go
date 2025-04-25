package config

import (
	"fmt"
	"sort"

	"github.com/goccy/go-yaml"
)

type Mode string

const (
	ModeFile Mode = "file"
	ModeLink Mode = "link"
	ModeDefault Mode = ModeFile
)

var (
	ErrVersionNotFound = fmt.Errorf("version not found")
	ErrModeNotFound    = fmt.Errorf("mode not found")
)

type Config struct {
	Version string               	`yaml:"version"`
	Mode		Mode               		`yaml:"mode"`
	Plans   map[string][]string	`yaml:"plans"`
	Groups  map[string]string    	`yaml:"groups"`
}

func New() *Config {
	return &Config{
		Version:	DotatoVersion(),
		Mode:			ModeDefault,
		Plans:   	map[string][]string{},
		Groups:  	map[string]string{}, // sample group
	}
}

func ParseConfig(data []byte) (*Config, error) {
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

func parseConfigFromStr(str string) (*Config, error) {
	return ParseConfig([]byte(str))
}

// For testing purpose
func (r Config) isEqual(other *Config) bool {
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
		if !compStrings(plan, other.Plans[key]) {
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

func compStrings(a []string, b []string) bool {
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
