package dotato

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-git/go-billy/v5"
	"github.com/msisdev/dotato/internal/ignore"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

var (
	ErrTooManySymlinks = fmt.Errorf("too many levels of symbolic links")
)

func (d Dotato) GetGroupIgnore(group string) (*ignore.Ignore, error) {
	if err := d.setConfig(); err != nil { return nil, err }

	dir := d.cdir.Copy()
	dir = append(dir, group)
	return readIgnoreRecur(d.fs, dir)
}

// Return true if file contents are equal
func (d Dotato) compareFile(a string, b string) (bool, error) {
	// Compare file sizes
	var size int64
	{
		sa, err := d.fs.Stat(a)
		if err != nil {
			return false, err
		}
		sb, err := d.fs.Stat(b)
		if err != nil {
			return false, err
		}
		if sa.Size() != sb.Size() {
			return false, nil
		}

		size = sa.Size()
	}

	// Open files
	var (
		fa billy.File
		fb billy.File
	)
	{
		var err error
		fa, err = d.fs.Open(a)
		if err != nil {
			return false, err
		}
		defer fa.Close()

		fb, err = d.fs.Open(b)
		if err != nil {
			return false, err
		}
		defer fb.Close()
	}

	// Compare file contents
	var (
		bufsiz = int64(4096)
		bufA = make([]byte, bufsiz)
		bufB = make([]byte, bufsiz)
		offset = int64(0)
	)
	for offset < size {
		// Read from file A
		_, err := fa.Read(bufA)
		if err != nil {
			if err == io.EOF {
				break
			}
			return false, err
		}

		// Read from file B
		n, err := fb.Read(bufB)
		if err != nil {
			if err == io.EOF {
				break
			}
			return false, err
		}

		// Compare buffers
		for i := 0; i < n; i++ {
			if bufA[i] != bufB[i] {
				return false, nil
			}
		}

		offset += int64(n)
	}

	return true, nil
}

// Returns corresponding path
func (d Dotato) DotToDtt(
	base gp.GardenPath,
	dot gp.GardenPath,
	group string,
) (path gp.GardenPath) {
	path = d.cdir.Copy()
	path = append(path, group)
	path = append(path, dot[len(base):]...)
	return
}

// Returns corresponding path
func (d Dotato) DttToDot(
	dtt gp.GardenPath,
	base gp.GardenPath,
) gp.GardenPath {
	path := base.Copy()
	for i := len(d.cdir)+1; i < len(dtt); i++ {
		path = append(path, dtt[i])
	}
	return path
}

// Overwrite is allowed
func (d Dotato) CreateAndCopyFile(src string, dst string) error {
	srcFile, err := d.fs.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := d.fs.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

const maxSymlinkDepth = 10	// prevent infinite loop

func evalSymlinksRecur(fs billy.Filesystem, path string, depth int) (string, error) {
	if depth > maxSymlinkDepth {
		return "", ErrTooManySymlinks
	}

	// Lstat
	fi, err := fs.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return path, nil
		}
		return "", err
	}

	// Not symlink ?
	if fi.Mode()&os.ModeSymlink == 0 {
		return path, nil
	}

	// Symlink, read the target
	// var target string
	target, err := fs.Readlink(path)
	if err != nil {
		return "", err
	}

	// Make absolute path
	resolved := target
	if !filepath.IsAbs(target) {
		dir := filepath.Dir(path)
		
		resolved = fs.Join(dir, target)

		// fs.Join() calls filepath.Clean() internally,
		// so we don't need to call it again.
		// 
		// resolved = filepath.Clean(resolved)
	}

	// Recursion
	return evalSymlinksRecur(fs, resolved, depth+1)
}


func (d Dotato) evalSymlinks(path string) (string, error) {
	return evalSymlinksRecur(d.fs, path, 0)
}
