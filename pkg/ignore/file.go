package ignore

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/pkg/gardenpath"
)

const (
	DefaultIgnoreFileName = "dotato.ignore"
)

var (
	ErrNotDir = fmt.Errorf("not a directory")
)

func CompileIgnoreFile(fs billy.Filesystem, filePath gardenpath.GardenPath) (*Rules, bool, error) {
	// Open file
	file, err := fs.Open(filePath.String())
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	defer file.Close()

	// Read file line by line
	buf := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		buf = append(buf, line)
	}
	
	return NewRules(buf...), true, nil
}

// CompileIgnoreFileRecur recursively compiles the ignore files
// that are in the subdirectories of the given path.
// The path must be a directory.
func CompileIgnoreFileRecur(fs billy.Filesystem, tree *RuleTree, dirPath gardenpath.GardenPath, name string) error {
	// Is path a directory?
	fi, err := fs.Stat(dirPath.String())
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return ErrNotDir
	}

	// Read ignore file from current directory
	rules, ok, err := CompileIgnoreFile(fs, append(dirPath, name))
	if err != nil {
		return err
	} else if ok {
		tree.Add(dirPath, rules)
	}

	// Get file infos
	fis, err := fs.ReadDir(dirPath.String())
	if err != nil {
		return err
	}

	// Find directories
	for _, fi := range fis {
		if !fi.IsDir() {
			continue
		}

		nextPath := append(dirPath, fi.Name())

		// Check if the directory is ignored
		if tree.Ignore(nextPath) {
			continue
		}

		// Recursively compile ignore files in subdirectories
		if err := CompileIgnoreFileRecur(fs, tree, nextPath, name); err != nil {
			return err
		}
	}

	return nil
}
