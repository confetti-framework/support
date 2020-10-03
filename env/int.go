package env

import (
	"github.com/lanvard/support"
	"os"
)

func Int(search string) int {
	env, OK := os.LookupEnv(search)
	if !OK {
		panic("env " + search + " not found")
	}

	return support.NewValue(env).Number()
}

func IntOr(search string, def int) int {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	result, err := support.NewValue(env).NumberE()
	if err != nil {
		return def
	}

	return result
}
