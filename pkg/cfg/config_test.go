package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	str := `
version: "0.0.1"

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
	cfg := &Config{
		Version: "0.0.1",
		Plans: map[string]GroupList{
			"all":      {},
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

	genCfg, err := NewConfigFromStr(str)
	assert.NoError(t, err)
	assert.True(t, genCfg.IsEqual(cfg), "Generated config should be equal to the expected config")
}
