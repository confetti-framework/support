package env

import (
	"os"
)

func Str(search string) string {
	env, _ := os.LookupEnv(search)

	return env
}

func StringOr(search string, def string) string {
	env, ok := os.LookupEnv(search)
	if !ok {
		return def
	}

	return env
}
