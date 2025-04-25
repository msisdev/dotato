package dotato

import (
	"os"
)

// Loop up in the env var or use default value
func useEnvOrDefault(envVar, defaultValue string) string {
	if val, ok := os.LookupEnv(envVar); ok {
		return val
	}
	return defaultValue
}
