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

// If file is not found, it will return empty rules.
func ReadIgnoreFile(fs billy.Filesystem, filepath string) (*Rules, bool, error) {
	// Open file
	file, err := fs.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false, nil
		} else {
			return nil, false, err
		}
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

// read the ignore file and add it to the tree
func ReadIgnoreFileAdd(fs billy.Filesystem, tree *RuleTree, dir gardenpath.GardenPath, filename string) error {
	// Read ignore file
	rules, ok, err := ReadIgnoreFile(fs, append(dir, filename).String())
	if err != nil {
		return err
	} else if ok {
		tree.AddRules(dir, rules)
	}

	return nil
}

// ReadIgnoreFileRecur recursively reads the ignore files
// that are in the subdirectories of the given path.
func ReadIgnoreFileRecur(fs billy.Filesystem, tree *RuleTree, dir gardenpath.GardenPath, filename string) error {
	// Read ignore file in dir
	rules, ok, err := ReadIgnoreFile(fs, append(dir, filename).String())
	if err != nil {
		return err
	} else if ok {
		tree.AddRules(dir, rules)
	}

	// Get file infos
	fis, err := fs.ReadDir(dir.String())
	if err != nil {
		return err
	}

	// Find directories
	for _, fi := range fis {
		if !fi.IsDir() {
			continue
		}

		nextPath := append(dir, fi.Name())

		// Check if the directory is ignored
		if tree.Ignore(nextPath) {
			continue
		}

		// Recursively read ignore files in subdirectories
		if err := ReadIgnoreFileRecur(fs, tree, nextPath, filename); err != nil {
			return err
		}
	}

	return nil
}
