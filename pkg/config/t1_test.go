package config

var t1s = `
version: "0.0.1"

mode: file

plans:
  default:  []
  fedora:   [kde, konsole, zsh]
  hyprland: [alacritty, fish, waybar]

groups:
  alacritty:  "~/.config/alacritty"
  bash:       "~"
  fish:       "~/.config/fish"
  kde:        "~"
  konsole:    "~/.config"
  waybar:     "~/.config/waybar"
  zsh:        "~"
`

var t1c = &Config{
	Version: "0.0.1",
	Mode:		ModeFile,
	Plans: map[string]GroupList{
		"default":  {},
		"fedora":   {"kde", "konsole", "zsh"},
		"hyprland": {"alacritty", "fish", "waybar"},
	},
	Groups: map[string]string{
		"alacritty": "~/.config/alacritty",
		"bash":      "~",
		"fish":      "~/.config/fish",
		"kde":       "~",
		"konsole":   "~/.config",
		"waybar":    "~/.config/waybar",
		"zsh":       "~",
	},
}
