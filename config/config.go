package config

import (
	"os"
)

func LoadEnv(key string, fallback string) string {
	// Optionally load env vars here

	env, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return env
}
