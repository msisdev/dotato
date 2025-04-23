package config

import "runtime/debug"

const (
	DotatoVersionUnknown = "unknown"
)

func GetDotatoVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return DotatoVersionUnknown
	}

	return info.Main.Version
}
