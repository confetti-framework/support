package env

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

// bootstrap .env file
func Bootstrap(envFile, envFileTest string) error {
	var file string

	if strings.HasSuffix(os.Args[0], ".test") {
		file = envFile
	} else {
		file = envFileTest
	}

	return godotenv.Load(file)
}
