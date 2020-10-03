package env

import (
	"github.com/lanvard/support"
	"os"
)

func Bool(search string) bool {
	env, OK := os.LookupEnv(search)
	if !OK {
		panic("Environment '" + search + "' not found")
	}

	return support.NewValue(env).Bool()
}

func BoolOr(search string, def bool) bool {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	result, err := support.NewValue(env).BoolE()
	if err != nil {
		return def
	}

	return result
}
