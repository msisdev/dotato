package fio

import (
	"os"

	"github.com/msisdev/dotato/pkg/cfg"
	"github.com/goccy/go-yaml"
)

// NewConfigFile creates a new config file.
func NewConfigFile(name string, perm os.FileMode, config *cfg.Config) error {
	// Create and open file
	file, err := os.OpenFile(
		name,
		os.O_EXCL|os.O_CREATE|os.O_WRONLY,	// fail if file exists
		perm,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	// Prepare yaml content
	yamlByte, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	// Write yaml content
	if _, err = file.Write(yamlByte); err != nil {
		return err
	}

	return nil
}

// ReadConfigFile reads a config file,
// parses it and
// returns a Config struct.
func ReadConfigFile(name string) (*cfg.Config, error) {
	// Open and read file
	yamlByte, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	// Parse
	var config cfg.Config
	if err := yaml.Unmarshal(yamlByte, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// WriteConfigFile writes a config file.
// The config file must exist.
// It overwrites the existing file.
func WriteConfigFile(name string, perm os.FileMode, config cfg.Config) error {
	// Open existing file
	file, err := os.OpenFile(
		name,
		os.O_WRONLY|os.O_TRUNC,	// overwrite file
		perm,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	// Prepare yaml content
	yaml, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	// Write yaml content
	_, err = file.Write(yaml)
	if err != nil {
		return err
	}

	return nil
}
