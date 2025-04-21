package cfg

import "fmt"

const sampleConfigFormat = `version: %s

plans:
  arch: [alacritty]

groups:
  alacritty: "~/.config/alacritty"
`

func GetSampleConfigStr() string {
	return fmt.Sprintf(sampleConfigFormat, getDotatoVersion())
}
