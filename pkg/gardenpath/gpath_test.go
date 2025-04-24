package gardenpath

import (
	"os"
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

	// PWD
	pwdStr := os.Getenv("PWD")
	if pwdStr == "" {
		panic("PWD environment variable is not set")
	}
	pwd, err := NewGardenPath(pwdStr)
	if err != nil {
		panic("Failed to create GardenPath from PWD: " + err.Error())
	}
	assert.Equal(t, "", pwd[0], "Expected first element of PWD to be empty string")
	assert.Equal(t, pwdStr, pwd.String(), "Expected PWD to match the string value of the environment variable")

	testcases := []Testcase{
		// Test empty path
		{"", nil},
		// Test pwd
		{".", GardenPath(pwd)},
		{"./", GardenPath(pwd)},
		{"./foo", GardenPath(append(pwd, "foo"))},
		{"foo", GardenPath(append(pwd, "foo"))},
		{"foo/", GardenPath(append(pwd, "foo"))},
		{"foo/bar", GardenPath(append(pwd, "foo", "bar"))},
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
			gp, err := NewGardenPath(tc.path)
			assert.NoError(t, err)
			assert.Equal(t, tc.gp, gp, "NewGardenPath(%s): expected %v, got %v", tc.path, tc.gp, gp)
		})
	}
}
