package config

import "fmt"

const configFormat = `# dotato.yaml
version: %s

# Mode is either file or link.
mode: file

# Plans are used to select multiple groups. No other purposes.
#  - There is no reserved plan name.
#  - Empty plan means all groups.
plans:
  all:          # empty plan means all groups
  mypc: [bash]  # select groups with list

# Groups are used to group files and tag them with path.
#  - There is no reserved group name nor resolver name.
#  - You can use multiple resolvers in the same group.
#  - You can use tilde or env var in a resolver.
#  - '%%FOO%%' is not handled. Use '$FOO' or '${FOO}' instead.
groups:
  bash:
    nux: "~"
    mac: "$HOME"
`

func GetExample() string {
	return fmt.Sprintf(configFormat, ConfigVersion)
}
