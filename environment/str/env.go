package str

import (
	_ "github.com/lanvard/support/environment"
	"os"
)

func Env(search string) string {
	env, OK := os.LookupEnv(search)
	if !OK {
		panic("Enviroment '" + search + "' not found")
	}

	return env
}

func EnvOr(search string, def string) string {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	return env
}
