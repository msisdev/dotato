package ignore

import (
	"bufio"
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/internal/lib/filesystem"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

// Ignore wraps RuleTree and implements file read.
type Ignore struct {
	fn string           // ignore file name
	fs billy.Filesystem // filesystem
	rt *RuleTree        // ignore rule tree
}

func _new(fs billy.Filesystem, base gp.GardenPath, filename string) *Ignore {
	i := Ignore{
		fn: filename,
		fs: fs,
		rt: newRuleTreeFromDir(base),
	}

	// Read the base ignore file
	// _, err := i.Read(base)
	// if err != nil {
	// 	panic(err)
	// }

	return &i
}

// Create an Ignore instance.
func New(base gp.GardenPath, ignoreFileName string) *Ignore {
	return _new(filesystem.NewOSFS(), base, ignoreFileName)
}

func NewWithFS(fs billy.Filesystem, base gp.GardenPath, ignoreFileName string) *Ignore {
	return _new(fs, base, ignoreFileName)
}

// Read the ignore file in the given directory.
func (i Ignore) Read(dir gp.GardenPath) (bool, error) {
	// Open file
	file, err := i.fs.Open(append(dir, i.fn).Abs())
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	defer file.Close()

	// Read file
	var buf []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		buf = append(buf, line)
	}

	i.rt.Append(dir, newRules(buf...))
	return true, nil
}

// Read the ignore file in the given directory and all its subdirectories.
func (i Ignore) ReadRecur(dir gp.GardenPath) error {
	// Read ignore file in dir
	_, err := i.Read(dir)
	if err != nil {
		return err
	}

	// Get file infos
	fis, err := i.fs.ReadDir(dir.Abs())
	if err != nil {
		return err
	}

	// Find directories
	for _, fi := range fis {
		if !fi.IsDir() {
			continue
		}

		// Read ignore file in subdir
		if err := i.ReadRecur(append(dir, fi.Name())); err != nil {
			return err
		}
	}

	return nil
}

func (i Ignore) IsIgnored(path gp.GardenPath) bool {
	return i.rt.IsIgnored(path)
}

func (i Ignore) IsIgnoredWithBaseDir(baseDir gp.GardenPath, path gp.GardenPath) bool {
	return i.rt.IsIgnoredWithBaseDir(baseDir, path)
}
