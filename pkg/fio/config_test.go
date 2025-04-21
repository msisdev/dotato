package fio

import (
	"fmt"
	"os"
	"testing"

	"github.com/msisdev/dotato/pkg/cfg"
)

const (
	filename = "config_test.yaml"
	fileperm     = 0644
)
// 
func openFile(name string) (*os.File, error) {
	file, err := os.OpenFile(
		name,
		os.O_RDWR,
		fileperm,
	)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func deleteFile(f *os.File) error {
	if f == nil {
		return fmt.Errorf("file is null")
	}

	name := f.Name()

	if err := f.Close(); err != nil {
		return err
	}

	if err := os.Remove(name); err != nil {
		return err
	}

	return nil
}

func TestNewConfigFile(t *testing.T) {
	expected := cfg.NewConfig()

	// Create a new config file
	if err := NewConfigFile(filename, fileperm, expected); err != nil {
		t.Fatalf("failed to create config file: %v", err)
	}

	// Read the same config file
	actual, err := ReadConfigFile(filename)
	if err != nil {
		t.Fatalf("failed to read config file: %v", err)
	}

	// Compare the Config structs
	if !expected.IsEqual(actual) {
		t.Fatalf("expected %v, got %v", expected, actual)
	}

	// Clean up
	file, err := openFile(filename)
	if err != nil {
		t.Fatalf("failed to open config file: %v", err)
	}
	deleteFile(file)	
}
