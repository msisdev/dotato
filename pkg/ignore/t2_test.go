package ignore

var t2e = []Entry{
	{ "/.foo", IsFile, Ignored },
	{ "/.bar", IsFile, NotIgnored },
	{ "/.baz", IsFile, NotIgnored },
	{ "/one", IsDir, NotIgnored },
	{ "/one/two", IsFile, NotIgnored },
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
