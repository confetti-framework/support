package boo

import (
	_ "github.com/lanvard/support/environment"
	"os"
)

func Env(search string) bool {
	env, OK := os.LookupEnv(search)
	if !OK {
		panic("Enviroment '" + search + "' not found")
	}

	return env == "true"
}

func EnvOr(search string, def bool) bool {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	return env == "true"
}
