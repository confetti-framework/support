package lang

import (
	_ "github.com/lanvard/support/environment"
	"golang.org/x/text/language"
	"os"
)

func Env(search string) language.Tag {
	env, OK := os.LookupEnv(search)
	if !OK {
		panic("Enviroment '" + search + "' not found")
	}

	return language.Make(env)
}

func EnvOr(search string, def language.Tag) language.Tag {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	return language.Make(env)
}
