package cfg

import "runtime/debug"

const (
	DotatoVersionUnknown = "unknown"
)

func getDotatoVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return DotatoVersionUnknown
	}

	return info.Main.Version
}
