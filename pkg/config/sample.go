package config

import "fmt"

const sampleConfigFormat = `version: %s

mode: file # file or link

plans:
  all: # empty plan means all groups
  # arch: [home] # select groups with list

groups:
  home: "~" # base directory of each group
  # bash: "$HOME" # you may use env vars
`

func GetSampleConfigStr() string {
	return fmt.Sprintf(sampleConfigFormat, DotatoVersion())
}
