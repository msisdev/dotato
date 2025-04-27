package ignore

var testcase2Files = []FileEntry{
	{"/.foo", IsFile, Ignored},
	{"/.bar", IsFile, NotIgnored},
	{"/.baz", IsFile, NotIgnored},
	{"/one", IsDir, NotIgnored},
	{"/one/two/the three", IsFile, NotIgnored},
}

// This rules start from `/` directory
var testcase2Rule = &ruleNode{
	rules: newRules(
		".*",
		"!.bar",
		"!.baz",
	),
	dirs: map[string]*ruleNode{},
}
