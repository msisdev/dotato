package ignore

import (
	"runtime"
	"testing"

	gp "github.com/msisdev/dotato/pkg/gardenpath"
	"github.com/stretchr/testify/assert"
)

func testRuleTree(t *testing.T, rt *RuleTree, fs []FileEntry) {
	for _, f := range fs {
		// garden path
		path, err := gp.New(f.path)
		assert.NoError(t, err)

		// test IsIgnored
		ignored := rt.IsIgnored(path)
		assert.Equal(t, f.isIgnored, ignored, "path: %s, isIgnored: %v", f.path, f.isIgnored)

		// test IsIgnoredWithBase
		remotePath := append(gp.GardenPath{"dummy"}, path...)
		ignored = rt.IsIgnoredWithBase(rt.base+1, remotePath)
		assert.Equal(t, f.isIgnored, ignored, "path: %s", f.path)
	}
}

func TestRuleTree1_Base0(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping Linux path test on non-Linux OS")
		return
	}

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
	if runtime.GOOS == "windows" {
		t.Skip("Skipping Linux path test on non-Linux OS")
		return
	}

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
	if runtime.GOOS == "windows" {
		t.Skip("Skipping Linux path test on non-Linux OS")
		return
	}

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
	if runtime.GOOS == "windows" {
		t.Skip("Skipping Linux path test on non-Linux OS")
		return
	}

	rt := &RuleTree{
		base: 0,
		head: testcase2Rule,
	}

	testRuleTree(t, rt, testcase2Files)
}

func TestRuleTree3(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping Linux path test on non-Linux OS")
		return
	}

	dir, err := gp.New("/home/user/.config/alacritty")
	assert.NoError(t, err)

	rt := &RuleTree{
		base: GetBaseFrom(dir),
		head: &ruleNode{
			rules: newRules(
				"*",
				"!alacritty.yml",
				"!themes/acme.toml",
			),
			dirs: map[string]*ruleNode{},
		},
	}

	{
		raw := "/home/user/.config/alacritty"
		expected := NotIgnored
		path, err := gp.New(raw)
		assert.NoError(t, err)

		ignored := rt.IsIgnored(path)
		assert.Equal(t, expected, ignored, "path: %s, isIgnored: %v", raw, expected)
	}

	{
		raw := "/home/user/.config/alacritty/alacritty.yml"
		expected := NotIgnored
		path, err := gp.New(raw)
		assert.NoError(t, err)
		ignored := rt.IsIgnored(path)
		assert.Equal(t, expected, ignored, "path: %s, isIgnored: %v", raw, expected)
	}

	{
		raw := "/home/user/.config/alacritty/themes"
		expected := Ignored
		path, err := gp.New(raw)
		assert.NoError(t, err)
		ignored := rt.IsIgnored(path)
		assert.Equal(t, expected, ignored, "path: %s, isIgnored: %v", raw, expected)
	}

	{
		raw := "/home/user/.config/alacritty/themes/acme.toml"
		expected := NotIgnored
		path, err := gp.New(raw)
		assert.NoError(t, err)
		ignored := rt.IsIgnored(path)
		assert.Equal(t, expected, ignored, "path: %s, isIgnored: %v", raw, expected)
	}

	{
		raw := "/home/user/.config/alacritty/themes/afterglow.toml"
		expected := Ignored
		path, err := gp.New(raw)
		assert.NoError(t, err)
		ignored := rt.IsIgnored(path)
		assert.Equal(t, expected, ignored, "path: %s, isIgnored: %v", raw, expected)
	}
}
