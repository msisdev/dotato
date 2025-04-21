package cfg

import "github.com/goccy/go-yaml"

type Config struct {
	Version	string								`yaml:"version"`
	Plans 	map[string]GroupList	`yaml:"plans"`
	Groups	map[string]string			`yaml:"groups"`
}

func NewConfig() *Config {
	return &Config{
		Version:	getDotatoVersion(),
		Plans: 		map[string]GroupList{},
		Groups:		map[string]string{},	// sample group
	}
}

func NewConfigFromStr(str string) (*Config, error) {
	var config Config
	if err := yaml.Unmarshal([]byte(str), &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// IsEqual can compare two Config objects.
// It is for testing purpose.
func (r Config) IsEqual(other *Config) bool {
	// Compare versions
	if r.Version != other.Version {
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
