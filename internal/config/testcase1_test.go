package config

var testcase1String = `
version: "0.0.1"

mode: file

plans:
  all:
  default:  []
  fedora:   [kde, konsole, zsh]
  hyprland: [alacritty, fish, waybar]

groups:
  alacritty:
    nux: "~/.config/alacritty"
  bash:
    nux: "~"
  fish:
    nux: "~/.config/fish"
  kde:
    nux: "~"
  konsole:
    nux: "~/.config"
  waybar:
    nux: "~/.config/waybar"
  zsh:
    nux: "~"
`

var testcase1Config = &Config{
	Version: "0.0.1",
	Mode:    ModeFile,
	Plans: map[string][]string{
		"all":      nil,
		"default":  {},
		"fedora":   {"kde", "konsole", "zsh"},
		"hyprland": {"alacritty", "fish", "waybar"},
	},
	// Groups: map[string]string{
	// 	"alacritty": "~/.config/alacritty",
	// 	"bash":      "~",
	// 	"fish":      "~/.config/fish",
	// 	"kde":       "~",
	// 	"konsole":   "~/.config",
	// 	"waybar":    "~/.config/waybar",
	// 	"zsh":       "~",
	// },
	Groups: map[string]map[string]string{
		"alacritty": {"nux": "~/.config/alacritty"},
		"bash":      {"nux": "~"},
		"fish":      {"nux": "~/.config/fish"},
		"kde":       {"nux": "~"},
		"konsole":   {"nux": "~/.config"},
		"waybar":    {"nux": "~/.config/waybar"},
		"zsh":       {"nux": "~"},
	},
}
