package dotato

import (
	"os"

	"github.com/go-git/go-billy/v5"
)

// Loop up in the env var or use default value
func useEnvOrDefault(envVar, defaultValue string) string {
	if val, ok := os.LookupEnv(envVar); ok {
		return val
	}
	return defaultValue
}

func createAndWriteFile(fs billy.Filesystem, path string, content []byte) error {
	file, err := fs.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	return err
}

func overwriteFile(fs billy.Filesystem, path string, content []byte) error {
	// Open file for writing
	file, err := fs.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	return err
}
