package config

import "fmt"

const configFormat = `version: %s

mode: file # file or link

plans:
  all: # empty plan means all groups
  # arch: [home] # select groups with list

groups:
  home:
    nux: "~" # base directory for home group in linux
  # macos: "$HOME" # you may use env vars
`

func GetExample() string {
	return fmt.Sprintf(configFormat, DotatoVersion())
}
