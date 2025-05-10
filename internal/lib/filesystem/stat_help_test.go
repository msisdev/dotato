package filesystem

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	gp "github.com/msisdev/dotato/pkg/gardenpath"
)

// Allows overwrite
func createFile(fs billy.Filesystem, path gp.GardenPath, content []byte) error {
	err := fs.MkdirAll(path.Parent().Abs(), 0755)
	if err != nil {
		return err
	}

	f, err := fs.Create(path.Abs())
	if err != nil {
		return err
	}

	_, err = f.Write(content)
	if err != nil {
		return err
	}

	return f.Close()
}

type FirstReq int
const (
	FirstReq_Empty FirstReq = iota
	FirstReq_File
	FirstReq_Link_Diff
	FirstReq_Link_Same
)

type SecondReq int
const (
	SecondReq_Empty SecondReq = iota
	SecondReq_File_NotEq
	SecondReq_File_Eq
	SecondReq_Link_Diff_NotEq
	SecondReq_Link_Diff_Eq
	SecondReq_Link_Same
)

var (
	randomPathFirst = gp.GardenPath{"", "random", "path", "first"}
	randomPathSecond = gp.GardenPath{"", "random", "path", "second"}
	fileContentEq = []byte("Hello, world!")
	fileContentNotEq = []byte("Hello, world! Alt")
)

func requestFS(
	first gp.GardenPath,
	firstReq FirstReq,
	second gp.GardenPath,
	secondReq SecondReq,
) billy.Filesystem {
	fs := memfs.New()

	switch firstReq {
	case FirstReq_Empty:
		// do nothing

	case FirstReq_File:
		err := createFile(fs, first, fileContentEq)
		if err != nil {
			panic(err)
		}

	case FirstReq_Link_Diff:
		err := createFile(fs, randomPathFirst, fileContentEq)
		if err != nil {
			panic(err)
		}
		err = fs.Symlink(randomPathFirst.Abs(), first.Abs())
		if err != nil {
			panic(err)
		}

	case FirstReq_Link_Same:
		err := fs.Symlink(second.Abs(), first.Abs())
		if err != nil {
			panic(err)
		}
	}

	switch secondReq {
	case SecondReq_Empty:
		// do nothing

	case SecondReq_File_NotEq:
		err := createFile(fs, second, fileContentNotEq)
		if err != nil {
			panic(err)
		}

	case SecondReq_File_Eq:
		err := createFile(fs, second, fileContentEq)
		if err != nil {
			panic(err)
		}

	case SecondReq_Link_Diff_NotEq:
		err := createFile(fs, randomPathSecond, fileContentNotEq)
		if err != nil {
			panic(err)
		}
		err = fs.Symlink(randomPathSecond.Abs(), second.Abs())
		if err != nil {
			panic(err)
		}

	case SecondReq_Link_Diff_Eq:
		err := createFile(fs, randomPathSecond, fileContentEq)
		if err != nil {
			panic(err)
		}
		err = fs.Symlink(randomPathSecond.Abs(), second.Abs())
		if err != nil {
			panic(err)
		}

	case SecondReq_Link_Same:
		err := fs.Symlink(first.Abs(), second.Abs())
		if err != nil {
			panic(err)
		}
	}

	return fs
}
