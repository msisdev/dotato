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
	base uint32    // Base index for the path
}

// Further algorithms will skip path[0:base] and
// start from path[base+1]
func newRuleTree(base uint32) *RuleTree {
	return &RuleTree{
		head: newNode(newRules()),
		base: base,
	}
}

// Provide a directory that you will use as a root of rule tree.
// This function will calculate the base index for you.
func newRuleTreeFromDir(dir gp.GardenPath) *RuleTree {
	return newRuleTree(uint32(len(dir)) - 1)
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

func (t *RuleTree) IsIgnored(path gp.GardenPath) bool {
	// Some edge cases
	if t.head == nil {
		t.head = newNode(newRules())
		return false
	}
	if len(path) == 0 {
		return false
	}

	node := t.head

	for idx, nextName := range path[t.base : len(path)-1] {
		if node.IsIgnored(path[idx:].Abs()) {
			return true
		}

		nextNode, ok := node.dirs[nextName]
		if !ok {
			return false
		}

		node = nextNode
	}

	return node.IsIgnored(path.Last())
}
