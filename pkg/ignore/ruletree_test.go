package ignore

import (
	"testing"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

type TreeHelper struct {
	tree *RuleTree
}

func (h TreeHelper) IsIgnored(path string) (bool, error) {
	// Make path
	gpath, err := gp.New(path)
	if err != nil {
		return false, err
	}

	// Test
	return h.tree.IsIgnored(gpath), nil
}

func (h TreeHelper) Test(t *testing.T, entries []FileEntry) {
	for _, entry := range entries {
		ignored, err := h.IsIgnored(entry.path)
		assert.NoError(t, err)
		assert.Equal(t, entry.isIgnored, ignored, "path: %s", entry.path)
	}
}

func TestRuleTree1_Base0(t *testing.T) {
	tree := &RuleTree{
		base: 0,
		head: &ruleNode{
			rules: newRules(),
			dirs: map[string]*ruleNode{
				"": {
					rules: newRules(),
					dirs: map[string]*ruleNode{
						"home": {
							rules: newRules(),
							dirs: map[string]*ruleNode{
								"user": testcase1Rule,
							},
						},
					},
				},
			},
		},
	}
	h := TreeHelper{tree}
	h.Test(t, testcase1Files)
}

func TestRuleTree1_Base1(t *testing.T) {
	tree := &RuleTree{
		base: 1,
		head: &ruleNode{
			rules: newRules(),
			dirs: map[string]*ruleNode{
				"home": {
					rules: newRules(),
					dirs: map[string]*ruleNode{
						"user": testcase1Rule,
					},
				},
			},
		},
	}

	h := TreeHelper{tree}
	h.Test(t, testcase1Files)
}

func TestRuleTree1_Base2(t *testing.T) {
	tree := &RuleTree{
		base: 2,
		head: &ruleNode{
			rules: newRules(),
			dirs: map[string]*ruleNode{
				"user": testcase1Rule,
			},
		},
	}

	h := TreeHelper{tree}
	h.Test(t, testcase1Files)
}

func TestRuleTree2_Base0(t *testing.T) {
	tree := &RuleTree{
		base: 0,
		head: testcase2Rule,
	}

	h := TreeHelper{tree}
	h.Test(t, testcase2Files)
}
