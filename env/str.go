package env

import (
	"os"
)

func Str(search string) string {
	env, _ := os.LookupEnv(search)

	return env
}

func StringOr(search string, def string) string {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	return env
}
