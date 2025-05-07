package ignore

import "fmt"

var testcase1Files = []FileEntry{
	{"/home/user", IsDir, NotIgnored},
	{"/home/user/.bash_history", IsFile, Ignored},
	{"/home/user/.bashrc", IsFile, NotIgnored},

	{"/home/user/.config", IsDir, NotIgnored},

	{"/home/user/.config/alacritty", IsDir, NotIgnored},
	{"/home/user/.config/alacritty/alacritty.yml", IsFile, NotIgnored},

	{"/home/user/.config/alacritty/themes", IsDir, NotIgnored},
	{"/home/user/.config/alacritty/themes/acme.toml", IsFile, Ignored},
	{"/home/user/.config/alacritty/themes/afterglow.toml", IsFile, Ignored},
	{"/home/user/.config/alacritty/themes/alabaster.toml", IsFile, NotIgnored},
	{"/home/user/.config/alacritty/themes/alabaster_dark.toml", IsFile, NotIgnored},
	{"/home/user/.config/alacritty/themes/ayu_dark.toml", IsFile, Ignored},
	{"/home/user/.config/alacritty/themes/ayu_light.toml", IsFile, Ignored},
	{"/home/user/.config/alacritty/themes/ayu_mirage.toml", IsFile, NotIgnored},

	{"/home/user/.config/fastfetch", IsDir, NotIgnored},
	{"/home/user/.config/fastfetch/config.jsonc", IsFile, NotIgnored},

	{"/home/user/.ssh", IsDir, NotIgnored},
	{"/home/user/.ssh/config", IsFile, NotIgnored},
	{"/home/user/.ssh/known_hosts", IsFile, Ignored},

	{"/home/user/readme/foo", IsDir, NotIgnored},
	{"/home/user/readme/foo/README.md", IsFile, Ignored},

	{"/home/user/readme/bar", IsDir, NotIgnored},
	{"/home/user/readme/bar/README.md", IsFile, Ignored},
}

// This rules start from `/home/user` directory
var testcase1Rule = &ruleNode{
	rules: newRules(
		"# comment",
		".bash_history",
	),
	dirs: map[string]*ruleNode{
		".config": {
			rules: newRules(),
			dirs: map[string]*ruleNode{
				"alacritty": {
					rules: newRules(),
					dirs: map[string]*ruleNode{
						"themes": {
							rules: newRules(
								"*",
								"!alabaster.toml",
								"!alabaster_dark.toml",
								"!ayu_mirage.toml",
							),
							dirs: map[string]*ruleNode{},
						},
					},
				},
			},
		},
		".ssh": {
			rules: newRules(
				"known_hosts",
			),
			dirs: map[string]*ruleNode{},
		},
		"readme": {
			rules: newRules(
				"**/README.md",
			),
			dirs: map[string]*ruleNode{},
		},
	},
}

// Equivalent ignore files of the above rules
var testcase1Ignore = []IgnoreEntry{
	{
		fmt.Sprintf("/home/user/%s", IgnoreFileName),
		[]string{".bash_history"},
	},
	{
		fmt.Sprintf("/home/user/.config/alacritty/themes/%s", IgnoreFileName),
		[]string{"*.toml", "!alabaster.toml", "!alabaster_dark.toml", "!ayu_mirage.toml"},
	},
	{
		fmt.Sprintf("/home/user/.ssh/%s", IgnoreFileName),
		[]string{"known_hosts"},
	},
	{
		fmt.Sprintf("/home/user/readme/%s", IgnoreFileName),
		[]string{"**/README.md"},
	},
}
