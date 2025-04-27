package ignore

import (
	"testing"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

func testRuleTree(t *testing.T, rt *RuleTree, fs []FileEntry) {
	for _, f := range fs {
		path, err := gp.New(f.path)
		assert.NoError(t, err)
		ignored := rt.IsIgnored(path)
		assert.Equal(t, f.isIgnored, ignored, "path: %s", f.path)
	}
}

func TestRuleTree1_Base0(t *testing.T) {
	rt := &RuleTree{
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

	testRuleTree(t, rt, testcase1Files)
}

func TestRuleTree1_Base1(t *testing.T) {
	rt := &RuleTree{
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

	testRuleTree(t, rt, testcase1Files)
}

func TestRuleTree1_Base2(t *testing.T) {
	rt := &RuleTree{
		base: 2,
		head: &ruleNode{
			rules: newRules(),
			dirs: map[string]*ruleNode{
				"user": testcase1Rule,
			},
		},
	}

	testRuleTree(t, rt, testcase1Files)
}

func TestRuleTree2_Base0(t *testing.T) {
	rt := &RuleTree{
		base: 0,
		head: testcase2Rule,
	}

	testRuleTree(t, rt, testcase2Files)
}
