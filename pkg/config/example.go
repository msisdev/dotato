package config

import "fmt"

const configFormat = `version: %s

mode: file # file or link

plans:
  all: # empty plan means all groups
  # my-pc: [home] # select groups with list

groups:
  bash:
    nux: "~" # base directory for home group in linux
    mac: "$HOME" # you may use env vars
`

func GetExample() string {
	return fmt.Sprintf(configFormat, Version1)
}
