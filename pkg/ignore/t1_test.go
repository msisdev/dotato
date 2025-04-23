package ignore

import "fmt"

var t1e = []Entry{
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
var t1r = &RuleNode{
	rules: NewRules(
		"# comment",
		".bash_history",
	),
	dirs: map[string]*RuleNode{
		".config": {
			rules: NewRules(),
			dirs: map[string]*RuleNode{
				"alacritty":	{
					rules: NewRules(),
					dirs: map[string]*RuleNode{
						"themes": {
							rules: NewRules(
								"*.toml",
								"!alabaster.toml",
								"!alabaster_dark.toml",
								"!ayu_mirage.toml",
							),
							dirs: map[string]*RuleNode{},
						},
					},
				},
			},
		},
		".ssh": {
			rules: NewRules(
				"known_hosts",
			),
			dirs: map[string]*RuleNode{},
		},
		"readme": {
			rules: NewRules(
				"**/README.md",
			),
			dirs: map[string]*RuleNode{},
		},
	},
}

// Equivalent ignore files of the above rules
var t1i = []IgnoreEntry{
	{
		fmt.Sprintf("/home/user/%s", DefaultIgnoreFileName),
		[]string{".bash_history"},
	},
	{
		fmt.Sprintf("/home/user/.config/alacritty/themes/%s", DefaultIgnoreFileName),
		[]string{"*.toml", "!alabaster.toml", "!alabaster_dark.toml", "!ayu_mirage.toml"},
	},
	{
		fmt.Sprintf("/home/user/.ssh/%s", DefaultIgnoreFileName),
		[]string{"known_hosts"},
	},
	{
		fmt.Sprintf("/home/user/readme/%s", DefaultIgnoreFileName),
		[]string{"**/README.md"},
	},
}
