package config

import (
	"io"
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/goccy/go-yaml"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

const (
	ConfigFileName = "dotato.yaml"
)

func Read(fs billy.Filesystem, filepath string) (*Config, bool, error) {
	// Open
	file, err := fs.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		}
		return nil, false, err
	}

	// Read
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, false, err
	}

	cfg, err := ParseConfig(b)

	return cfg, true, err
}

// ReadConfigRecur tries to find config file by
// walking up the directory tree.
// It returns the directory of the config file.
func ReadRecur(fs billy.Filesystem, dir gp.GardenPath, filename string) (*Config, gp.GardenPath, error) {
	if dir == nil {
		return nil, nil, nil
	}

	filepath := append(dir, filename)

	cfg, ok, err := Read(fs, filepath.String())
	if err != nil {
		return nil, nil, err
	}
	if ok {
		return cfg, dir, nil
	}

	return ReadRecur(fs, dir.Parent(), filename)
}

func Write(fs billy.Filesystem, filepath string, cfg *Config) error {
	// Open
	file, err := fs.OpenFile(
		filepath,
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
