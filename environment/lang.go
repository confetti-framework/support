package environment

import (
	"golang.org/x/text/language"
	"os"
)

func LangEnv(search string) language.Tag {
	env, OK := os.LookupEnv(search)
	if !OK {
		panic("Enviroment '" + search + "' not found")
	}

	return language.Make(env)
}

func LangEnvOr(search string, def language.Tag) language.Tag {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	return language.Make(env)
}
