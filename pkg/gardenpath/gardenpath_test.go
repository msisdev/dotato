package gardenpath

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func test(t *testing.T, path string, expected GardenPath, expectedAbs string) {
	gp, err := New(path)
	assert.NoError(t, err, "New(%s): %v", path, err)
	assert.Equal(t, expected, gp, "New(%s): expected %v, got %v", path, expected, gp)
	assert.Equal(t, expectedAbs, gp.Abs(), "New(%s): expected Abs() %s, got %s", path, expectedAbs, gp.Abs())
}

func TestGardenPathLinux(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping Linux path test on non-Linux OS")
		return
	}

	// Get WD
	wdStr, err := os.Getwd()
	if err != nil {
		panic("Failed to get current working directory: " + err.Error())
	}
	wdgp, err := New(wdStr)
	if err != nil {
		panic("Failed to create GardenPath from PWD: " + err.Error())
	}
	assert.Equal(t, "", wdgp[0])
	assert.Equal(t, wdStr, wdgp.Abs(), "Expected %s, got %s", wdStr, wdgp.Abs())

	test(t, "", nil, "")
	test(t, ".", wdgp, wdStr)
	test(t, "./", wdgp, wdStr)
	test(t, "./foo", append(wdgp.Copy(), "foo"), wdStr+"/foo")
	test(t, "foo", append(wdgp.Copy(), "foo"), wdStr+"/foo")
	test(t, "foo/", append(wdgp.Copy(), "foo"), wdStr+"/foo")
	test(t, "foo/bar", append(wdgp.Copy(), "foo", "bar"), wdStr+"/foo/bar")
	test(t, "/", GardenPath{""}, "/")
	test(t, "/home", GardenPath{"", "home"}, "/home")
	test(t, "/home/", GardenPath{"", "home"}, "/home")
	test(t, "/home/user", GardenPath{"", "home", "user"}, "/home/user")
	test(t, "~", GardenPath{"", "home", "user"}, "/home/user")
	test(t, "~/", GardenPath{"", "home", "user"}, "/home/user")
	test(t, "~/foo", GardenPath{"", "home", "user", "foo"}, "/home/user/foo")
}

func TestGardenPathWindows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping Windows path test on non-Windows OS")
		return
	}

	// Get WD
	wdStr, err := os.Getwd()
	if err != nil {
		panic("Failed to get current working directory: " + err.Error())
	}
	wdgp, err := New(wdStr)
	if err != nil {
		panic("Failed to create GardenPath from PWD: " + err.Error())
	}
	assert.Equal(t, wdStr, wdgp.Abs(), "Expected %s, got %s", wdStr, wdgp.Abs())

	test(t, "", nil, "")
	test(t, "C:", wdgp, wdStr)
	test(t, ".", wdgp, wdStr)
	test(t, "C:\\", GardenPath{"C:"}, "C:\\")
	test(t, "C:\\home", GardenPath{"C:", "home"}, "C:\\home")
	test(t, "C:\\home\\", GardenPath{"C:", "home"}, "C:\\home")
	test(t, "C:\\home\\user", GardenPath{"C:", "home", "user"}, "C:\\home\\user")
}

func TestCopy(t *testing.T) {
	path := GardenPath{"foo", "bar", "baz"}
	path2 := path.Copy()
	assert.Equal(t, path, path2, "Expected %v, got %v", path, path2)
}
