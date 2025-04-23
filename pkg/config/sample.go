package config

import "fmt"

const sampleConfigFormat = `version: %s

mode: file

plans:
  desktop: [bash]

groups:
  bash: "~"
`

func GetSampleConfigStr() string {
	return fmt.Sprintf(sampleConfigFormat, GetDotatoVersion())
}
