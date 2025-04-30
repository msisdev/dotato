package ignore

var testcase2Files = []FileEntry{
	{"/.foo", IsFile, Ignored},
	{"/.bar", IsFile, NotIgnored},
	{"/.baz", IsFile, NotIgnored},
	{"/one", IsDir, Ignored},
	{"/one/two/the three", IsFile, Ignored},
}

// This rules start from `/` directory
var testcase2Rule = &ruleNode{
	rules: newRules(
		"**/*",
		"!.bar",
		"!.baz",
	),
	dirs: map[string]*ruleNode{},
}
