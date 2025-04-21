package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	list := GroupList{"a", "b", "c"}
	set := GroupSet{"a": true, "b": true, "c": true}

	assert.Equal(t, list.ToGroupSet(), set)
	assert.Equal(t, set.ToGroupList(), list)
	assert.True(t, list.IsEqual(GroupList{"a", "b", "c"}))
	assert.False(t, list.IsEqual(GroupList{"a", "b"}))
	assert.True(t, set.IsEqual(GroupSet{"a": true, "b": true, "c": true}))
	assert.False(t, set.IsEqual(GroupSet{"a": true, "b": true}))
}
