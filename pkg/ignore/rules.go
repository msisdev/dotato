package ignore

import (
	"github.com/sabhiram/go-gitignore"
)

// Rule is a wrapper of external gitignore package.
type Rules struct {
	i *ignore.GitIgnore
}

// Passing nothing will create an empty rule.
func NewRules(lines ...string) *Rules {
	return &Rules{
		i: ignore.CompileIgnoreLines(lines...),
	}
}

// Ignore rule works with relative path.
func (r Rules) Ignore(relPath string) bool {
	return r.i.MatchesPath(relPath)
}
