package gardenpath

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGardenPath(t *testing.T) {
	type Testcase struct {
		path 	string
		gp		GardenPath
	}

	// HOME
	os.Setenv("HOME", "/home/user")

	// WD
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
		{"", nil},
		// Test wd
		{".", wd},
		{"./", wd},
		{"./foo", append(wd, "foo")},
		{"foo", append(wd, "foo")},
		{"foo/", append(wd, "foo")},
		{"foo/bar", append(wd, "foo", "bar")},
		// Test root path
		{"/", GardenPath{""}},
		// Test normal paths
		{"/home", GardenPath{"", "home"}},
		{"/home/", GardenPath{"", "home"}},
		{"/home/user", GardenPath{"", "home", "user"}},
		// Test tilde
		{"~", GardenPath{"", "home", "user"}},
		{"~/", GardenPath{"", "home", "user"}},
		{"~/foo", GardenPath{"", "home", "user", "foo"}},
		// Test env vars
		{"$HOME", GardenPath{"", "home", "user"}},
		{"$HOME/", GardenPath{"", "home", "user"}},
		{"$HOME/foo", GardenPath{"", "home", "user", "foo"}},
		// Test env vars 2
		{"${HOME}", GardenPath{"", "home", "user"}},
		{"${HOME}/", GardenPath{"", "home", "user"}},
		{"${HOME}/foo", GardenPath{"", "home", "user", "foo"}},
	}

	for _, tc := range testcases {
		t.Run(tc.path, func(t *testing.T) {
			gp, err := New(tc.path)
			assert.NoError(t, err)
			assert.Equal(t, tc.gp, gp, "New(%s): expected %v, got %v", tc.path, tc.gp, gp)
		})
	}
}

func TestWindowsPath(t *testing.T) {
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
