package config

import "fmt"

const sampleConfigFormat = `version: %s

mode: file # file or link are supported

plans:
  all: # empty plan means all groups
  # arch: [home] # select groups with list

groups:
  home: "~" # 
  # bash: "~"
`

func GetSampleConfigStr() string {
	return fmt.Sprintf(sampleConfigFormat, GetDotatoVersion())
}
