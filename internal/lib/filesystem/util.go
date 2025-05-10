package filesystem

import (
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-git/go-billy/v6"
)

func GetRootDir() string {
	if runtime.GOOS == "windows" {
		if drive := os.Getenv("SystemDrive"); drive != "" {
			return drive + "\\"
		}
		panic("Oops! SystemDrive env var is not set. Please set it with your current drive. E.g. 'C:'")
	}
	return "/"
}

// Return true if file contents are equal
func IsFileContentEqual(fs billy.Filesystem, a, b string) (bool, error) {
	// Compare file sizes
	var size int64
	{
		sa, err := fs.Stat(a)
		if err != nil {
			return false, err
		}
		sb, err := fs.Stat(b)
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
		fa, err = fs.Open(a)
		if err != nil {
			return false, err
		}
		defer fa.Close()

		fb, err = fs.Open(b)
		if err != nil {
			return false, err
		}
		defer fb.Close()
	}

	// Compare file contents
	var (
		bufsiz = int64(1024)
		bufA   = make([]byte, bufsiz)
		bufB   = make([]byte, bufsiz)
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

// Overwrite is allowed
func CreateAndCopyFile(fs billy.Filesystem, src, dst string) error {
	srcFile, err := fs.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := fs.Create(dst)
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

func evalSymlinks(fs billy.Filesystem, path string, depth int) (string, error) {
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
	return evalSymlinks(fs, resolved, depth+1)
}

func EvalSymlinks(fs billy.Filesystem, path string) (string, error) {
	return evalSymlinks(fs, path, 0)
}
