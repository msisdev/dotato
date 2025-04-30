package gardenpath

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGardenPathLinux(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping Linux path test on non-Linux OS")
		return
	}

	type Testcase struct {
		path 	string
		gp		GardenPath
		abs		string
	}

	// HOME
	os.Setenv("HOME", "/home/user")

	// Get WD
	wdStr, err := os.Getwd()
	if err != nil {
		panic("Failed to get current working directory: " + err.Error())
	}
	wd, err := New(wdStr)
	if err != nil {
		panic("Failed to create GardenPath from PWD: " + err.Error())
	}
	assert.Equal(t, "", wd[0])
	assert.Equal(t, wdStr, wd.Abs(), "Expected %s, got %s", wdStr, wd.Abs())

	testcases := []Testcase{
		// Test empty path
		{"", nil, ""},
		// Test wd
		{".", wd, wdStr},
		{"./", wd, wdStr},
		{"./foo", append(wd, "foo"), filepath.Join(wdStr, "foo")},
		{"foo", append(wd, "foo"), filepath.Join(wdStr, "foo")},
		{"foo/", append(wd, "foo"), filepath.Join(wdStr, "foo")},
		{"foo/bar", append(wd, "foo", "bar"), filepath.Join(wdStr, "foo", "bar")},
		// Test root path
		{"/", GardenPath{""}, "/"},
		// Test normal paths
		{"/home", GardenPath{"", "home"}, "/home"},
		{"/home/", GardenPath{"", "home"}, "/home"},
		{"/home/user", GardenPath{"", "home", "user"}, "/home/user"},
		// Test tilde
		{"~", GardenPath{"", "home", "user"}, "/home/user"},
		{"~/", GardenPath{"", "home", "user"}, "/home/user"},
		{"~/foo", GardenPath{"", "home", "user", "foo"}, "/home/user/foo"},
		// Test env vars
		// {"$HOME", GardenPath{"", "home", "user"}, "/home/user"},
		// {"$HOME/", GardenPath{"", "home", "user"}, "/home/user"},
		// {"$HOME/foo", GardenPath{"", "home", "user", "foo"}, "/home/user/foo"},
		// Test env vars 2
		// {"${HOME}", GardenPath{"", "home", "user"}, "/home/user"},
		// {"${HOME}/", GardenPath{"", "home", "user"}, "/home/user"},
		// {"${HOME}/foo", GardenPath{"", "home", "user", "foo"}, "/home/user/foo"},
	}

	for _, tc := range testcases {
		t.Run(tc.path, func(t *testing.T) {
			gp, err := New(tc.path)
			assert.NoError(t, err)
			assert.Equal(t, tc.gp, gp, "New(%s): expected %v, got %v", tc.path, tc.gp, gp)
			assert.Equal(t, tc.abs, gp.Abs(), "Abs(%s): expected %s, got %s", tc.path, tc.abs, gp.Abs())
		})
	}
}

func TestGardenPathWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping Windows path test on non-Windows OS")
		return
	}

	type Testcase struct {
		path 	string
		gp		GardenPath
	}

	testcases := []Testcase{
		// Test empty path
		{"", nil},
		// Test root path
		{"C:\\", GardenPath{""}},
		// Test normal paths
		{"C:\\home", GardenPath{"", "home"}},
		{"C:\\home\\", GardenPath{"", "home"}},
		{"C:\\home\\user", GardenPath{"", "home", "user"}},
	}

	for _, tc := range testcases {
		t.Run(tc.path, func(t *testing.T) {
			gp, err := New(tc.path)
			assert.NoError(t, err)
			assert.Equal(t, tc.gp, gp, "New(%s): expected %v, got %v", tc.path, tc.gp, gp)
		})
	}
}
