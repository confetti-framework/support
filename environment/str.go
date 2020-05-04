package environment

import (
	"os"
)

func StrEnv(search string) string {
	env, _ := os.LookupEnv(search)

	return env
}

func StrEnvOr(search string, def string) string {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	return env
}
