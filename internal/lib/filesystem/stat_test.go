package filesystem

import (
	"os"
	"runtime"
	"testing"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

func getGardenPathFirstEl() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("SystemDrive")
	}
	return ""
}

func TestNewPathStat(t *testing.T) {
	var el = getGardenPathFirstEl()

	path := gp.GardenPath{el, "home", "user", ".bashrc"}
	real := gp.GardenPath{el, "home", "user", "Documents", "dotato", "bash", ".bashrc"}

	// Path: link / Real: file
	{
		fs := requestFS(real, FirstReq_File, path, SecondReq_Link_Same)
		stat, err := NewPathStat(fs, path)
		assert.NoError(t, err)
		assert.Equal(t, path, stat.Path)
		assert.Equal(t, real, stat.Target)
		assert.Equal(t, real, stat.Real)
		assert.Equal(t, false, stat.IsFile)
		assert.Equal(t, true, stat.Exists)
	}

	// Path: file
	{
		fs := requestFS(real, FirstReq_File, path, SecondReq_File_Eq)
		stat, err := NewPathStat(fs, path)
		assert.NoError(t, err)
		assert.Equal(t, path, stat.Path)
		assert.Equal(t, path, stat.Target)
		assert.Equal(t, path, stat.Real)
		assert.Equal(t, true, stat.IsFile)
		assert.Equal(t, true, stat.Exists)
	}
}
