package config

import (
	"io"
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/goccy/go-yaml"
)

const (
	DefaultConfigFileName = "dotato.yaml"
)

func ReadFile(fs billy.Filesystem, name string) (*Config, error) {
	// Open
	file, err := fs.Open(name)
	if err != nil {
		return nil, err
	}

	// Read
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return NewFromByte(b)
}

func WriteFile(fs billy.Filesystem, name string, cfg *Config) error {
	// Open
	file, err := fs.OpenFile(
		name,
		os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_EXCL,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	// Prepare yaml
	yamlByte, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	// Write
	if _, err = file.Write(yamlByte); err != nil {
		return err
	}

	return nil
}
