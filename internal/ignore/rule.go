package ignore

import (
	gp "github.com/msisdev/dotato/pkg/gardenpath"
	ignore "github.com/sabhiram/go-gitignore"
)

// A wrapper of external gitignore package.
type Rules struct {
	i *ignore.GitIgnore
}

// Pass nothing to create an empty rule.
func newRules(lines ...string) *Rules {
	return &Rules{
		i: ignore.CompileIgnoreLines(lines...),
	}
}

// Ignore rule works with relative path.
func (r Rules) IsIgnored(relPath string) bool {
	return r.i.MatchesPath(relPath)
}

/////////////////////////////////////////////////

type ruleNode struct {
	rules *Rules               // Rules for the current directory
	dirs  map[string]*ruleNode // Subdirectories
}

func newNode(rules *Rules) *ruleNode {
	return &ruleNode{
		rules: rules,
		dirs:  make(map[string]*ruleNode),
	}
}

func (n *ruleNode) Append(dir string, node *ruleNode) {
	if _, ok := n.dirs[dir]; ok {
		return
	}
	n.dirs[dir] = node
}

func (n *ruleNode) Set(rules *Rules) {
	n.rules = rules
}

func (n *ruleNode) IsIgnored(relPath string) bool {
	if n.rules == nil {
		return false
	}
	return n.rules.IsIgnored(relPath)
}

/////////////////////////////////////////////////

type RuleTree struct {
	head *ruleNode // head is a dummy node
	base int		   // Base index for the path
}

// Further algorithms will skip path[0:base] and
// start from path[base+1]
func newRuleTree(base int) *RuleTree {
	return &RuleTree{
		head: newNode(newRules()),
		base: base,
	}
}

// Provide a directory that you will use as a root of rule tree.
// This function will calculate the base index for you.
func newRuleTreeFromDir(dir gp.GardenPath) *RuleTree {
	return newRuleTree(GetBaseFrom(dir))
}

// Add rules to the tree at the given dir.
//
// This function overwrites the existing rules.
//
// If dir is nil, it will do nothing.
func (t *RuleTree) Append(dir gp.GardenPath, rules *Rules) {
	// Some edge cases
	if dir == nil {
		return
	}
	if t.head == nil {
		t.head = newNode(newRules())
	}
	if GetBaseFrom(dir) < t.base {
		return
	}

	node := t.head
	for _, nextDir := range dir[t.base:] {
		nextNode, ok := node.dirs[nextDir]
		if !ok {
			// Append a new node
			nextNode = newNode(newRules())
			node.Append(nextDir, nextNode)
		}

		node = nextNode
	}

	// Set rules
	node.Set(rules)
}

func (t RuleTree) IsIgnored(path gp.GardenPath) bool {
	return t.IsIgnoredWithBase(t.base, path)
}

// Use base from the given dir.
func (t RuleTree) IsIgnoredWithBaseDir(baseDir gp.GardenPath, path gp.GardenPath) bool {
	return t.IsIgnoredWithBase(GetBaseFrom(baseDir), path)
}

// 
func (t RuleTree) IsIgnoredWithBase(base int, path gp.GardenPath) bool {
	// Some edge cases
	if t.head == nil {
		return false
	}
	if path == nil {
		return false
	}
	if base >= len(path) {
		// this path is same with or above the base
		return false
	}

	node := t.head
	
	if parent := path.Parent(); base < len(parent) {
		for i, nextName := range parent[base:] {
			if node.IsIgnored(path[i:].Abs()) {
				return true
			}
	
			nextNode, ok := node.dirs[nextName]
			if !ok {
				return false
			}
			node = nextNode
		}
	}

	return node.IsIgnored(path.Last())
}

func (t *RuleTree) SetBase(base int) {
	t.base = base
}

func (t *RuleTree) SetBaseFromDir(dir gp.GardenPath) {
	t.SetBase(GetBaseFrom(dir))
}

func GetBaseFrom(dir gp.GardenPath) int {
	return len(dir)
}
