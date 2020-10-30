package env

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func init() {
	Bootstrap()
}

// bootstrap .env file
func Bootstrap() {
	var file string
	root, _ := os.Getwd()

	if strings.HasSuffix(os.Args[0], ".test") {
		file = root + "/.env.testing"
	} else {
		file = root + "/.env"
	}

	_ = godotenv.Load(file)
}
