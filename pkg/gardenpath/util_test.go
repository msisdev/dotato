package gardenpath

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceTilde(t *testing.T) {
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
			p, err := replaceTilde(tc[0])
			assert.NoError(t, err)
			assert.Equal(t, tc[1], p)
		})
	}
}
