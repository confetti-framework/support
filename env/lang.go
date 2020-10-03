package env

import (
	"golang.org/x/text/language"
	"os"
)

func Lang(search string) language.Tag {
	env, OK := os.LookupEnv(search)
	if !OK {
		panic("Enviroment '" + search + "' not found")
	}

	return language.Make(env)
}

func LangOr(search string, def language.Tag) language.Tag {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	return language.Make(env)
}
