package ignore

import (
	"github.com/msisdev/dotato/pkg/gardenpath"
)

////////////////////////////////////////////////

type RuleNode struct {
	rules	*Rules								// Rules for the current directory
	dirs	map[string]*RuleNode	// Subdirectories
}

func NewRuleNode(rules *Rules) *RuleNode {
	return &RuleNode{
		rules: rules,
		dirs:  make(map[string]*RuleNode),
	}
}

func (n *RuleNode) AddRules(dir string, node *RuleNode) {
	if _, ok := n.dirs[dir]; ok {
		return
	}
	n.dirs[dir] = node
}

func (n *RuleNode) Ignore(relPath string) bool {
	if n.rules == nil {
		return false
	}
	return n.rules.Ignore(relPath)
}

////////////////////////////////////////////////

type RuleTree struct {
	root	*RuleNode
	base	uint32		// Base index for the path
}

// Further algorithms will skip path[0:base] and
// start from path[base+1]
func NewRuleTree(base uint32) *RuleTree {
	return &RuleTree{
		root: NewRuleNode(NewRules()),
		base: base,
	}
}

// This function will calculate the base index
// from the given base directory.
func NewRuleTreeFromBase(base gardenpath.GardenPath) *RuleTree {
	return &RuleTree{
		root: NewRuleNode(NewRules()),
		base: uint32(len(base))-1,
	}
}

// This function may overwrite the rules.
func (t *RuleTree) AddRules(path gardenpath.GardenPath, rules *Rules) {
	// Some edge cases
	if t.root == nil {
		t.root = NewRuleNode(NewRules())
	}

	// Traverse tree with given path
	// If given path is empty, rules will be added to the root node.
	node := t.root
	base := max(t.base, 1)
	for i := base; i < uint32(len(path)); i++ {
		dir := path[i]
		if next, ok := node.dirs[dir]; !ok {
			next = NewRuleNode(NewRules())
			node.AddRules(dir, next)
			node = next
		} else {
			node = next
		}
	}

	// Set rules
	node.rules = rules
}

func (t *RuleTree) Ignore(path gardenpath.GardenPath) bool {
	// Some edge cases
	if t.root == nil {
		t.root = NewRuleNode(NewRules())
		return false
	}
	if len(path) == 0 {
		return false
	}

	// From root to the parent node of the path,
	// check if the path is ignored on every node.
	node := t.root
	base := t.base + 1
	parent := uint32(len(path) - 1)

	// Always apply root node rules 
	if node.Ignore(path.String()) {
		return true
	}

	// Find next node and apply rules
	for idx := base; idx < parent; idx++ {
		dir := path[idx]

		// Go to next node
		if next, ok := node.dirs[dir]; !ok {
			return false
		} else {
			node = next
		}

		// Check rules
		if node.Ignore(path[idx:].String()) {
			return true
		}
	}

	return false
}
