package ignore

import (
	"testing"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

type TreeHelper struct {
	tree *RuleTree
}

func (h TreeHelper) Ignore(path string) (bool, error) {
	// Make path
	gpath, err := gp.New(path)
	if err != nil {
		return false, err
	}

	// Test
	return h.tree.Ignore(gpath), nil
}

func (h TreeHelper) Test(t *testing.T, entries []Entry) {
	for _, entry := range entries {
		ignored, err := h.Ignore(entry.path)
		assert.NoError(t, err)
		assert.Equal(t, entry.isIgnored, ignored, "path: %s", entry.path)
	}
}

func TestRuleTree1_Base0(t *testing.T) {
	tree := &RuleTree{
		base: 0,
		root: &RuleNode{
			rules: NewRules(),
			dirs: map[string]*RuleNode{
				"home": {
					rules: NewRules(),
					dirs: map[string]*RuleNode{
						"user": t1r,
					},
				},
			},
		},
	}

	h := TreeHelper{tree}
	h.Test(t, t1e)
}

func TestRuleTree1_Base1(t *testing.T) {
	tree := &RuleTree{
		base: 1,
		root: &RuleNode{
			rules: NewRules(),
			dirs: map[string]*RuleNode{
				"user": t1r,
			},
		},
	}

	h := TreeHelper{tree}
	h.Test(t, t1e)
}

func TestRuleTree1_Base2(t *testing.T) {
	tree := &RuleTree{
		base: 2,
		root: t1r,
	}

	h := TreeHelper{tree}
	h.Test(t, t1e)
}

func TestRuleTree2_Base0(t *testing.T) {
	tree := &RuleTree{
		base: 0,
		root: t2r,
	}

	h := TreeHelper{tree}
	h.Test(t, t2e)
}
