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

	os.Setenv("HOME", "/home/user")

	testcases := []Testcase{
		{"/", GardenPath{""}},
		{"/home", GardenPath{"", "home"}},
		{"/home/", GardenPath{"", "home"}},
		{"/home/user", GardenPath{"", "home", "user"}},
		{"~", GardenPath{"", "home", "user"}},
		{"~/", GardenPath{"", "home", "user"}},
		{"~/foo", GardenPath{"", "home", "user", "foo"}},
		{"$HOME", GardenPath{"", "home", "user"}},
		{"$HOME/", GardenPath{"", "home", "user"}},
		{"$HOME/foo", GardenPath{"", "home", "user", "foo"}},
		{"${HOME}", GardenPath{"", "home", "user"}},
		{"${HOME}/", GardenPath{"", "home", "user"}},
		{"${HOME}/foo", GardenPath{"", "home", "user", "foo"}},
	}

	for _, tc := range testcases {
		t.Run(tc.path, func(t *testing.T) {
			gp, err := New(tc.path)
			assert.NoError(t, err)
			assert.Equal(t, tc.gp, gp, "NewGardenPath(%s): expected %v, got %v", tc.path, tc.gp, gp)
		})
	}
}
