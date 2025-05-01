package cli

import "runtime/debug"

const (
	DotatoVersionUnknown = "unknown"
)

func dotatoVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return DotatoVersionUnknown
	}

	return info.Main.Version
}
