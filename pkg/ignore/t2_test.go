package ignore

var t2e = []FileEntry{
	{ "/.foo", IsFile, Ignored },
	{ "/.bar", IsFile, NotIgnored },
	{ "/.baz", IsFile, NotIgnored },
	{ "/one", IsDir, NotIgnored },
	{ "/one/two/the three", IsFile, NotIgnored },
}

// This rules start from `/` directory
var t2r = &RuleNode{
	rules: NewRules(
		".*",
		"!.bar",
		"!.baz",
	),
	dirs: map[string]*RuleNode{},
}

