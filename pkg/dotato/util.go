package dotato

import (
	"io"
	
	"github.com/go-git/go-billy/v5"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/msisdev/dotato/pkg/ignore"
)

func (d Dotato) GetGroupIgnore(group string) (*ignore.Ignore, error) {
	if err := d.setConfig(); err != nil { return nil, err }

	return readIgnoreRecur(d.fs, append(d.cdir, group))
}

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
	path = append(d.cdir, group)
	path = append(path, dot[len(base):]...)
	return
}

// Returns corresponding path
func (d Dotato) DttToDot(
	dtt gp.GardenPath,
	base gp.GardenPath,
) gp.GardenPath {
	return append(base, dtt[len(d.cdir)+1:]...)
}

func (d Dotato) CopyFile(src string, dst string) error {
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
