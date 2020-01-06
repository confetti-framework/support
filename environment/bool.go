package environment

import (
	"os"
)

func BoolEnv(search string) bool {
	env, OK := os.LookupEnv(search)
	if !OK {
		panic("Enviroment '" + search + "' not found")
	}

	return env == "true"
}

func BoolEnvOr(search string, def bool) bool {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	return env == "true"
}
