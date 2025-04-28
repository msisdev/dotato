package gardenpath

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandTilde(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Tilde expansion is only supported on Linux")
	}

	os.Setenv("HOME", "/home/user")
	user := os.Getenv("USER")

	testcases := [][2]string{
		{"/", "/"},
		{"~", "/home/user"},
		{"~/", "/home/user/"},
		{"~/foo", "/home/user/foo"},
		{fmt.Sprintf("~%s", user), fmt.Sprintf("/home/%s", user)},
		{fmt.Sprintf("~%s/", user), fmt.Sprintf("/home/%s/", user)},
		{fmt.Sprintf("~%s/foo", user), fmt.Sprintf("/home/%s/foo", user)},
	}

	for _, tc := range testcases {
		t.Run(tc[0], func(t *testing.T) {
			p, err := expandTilde(tc[0])
			assert.NoError(t, err)
			assert.Equal(t, tc[1], p)
		})
	}
}

func TestExpandEnv(t *testing.T) {
	
}