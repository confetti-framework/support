package env

import (
	"github.com/joho/godotenv"
	"os"
)

func init() {
	Bootstrap()
}

// bootstrap .env file
func Bootstrap() {
	root, _ := os.Getwd()
	_ = godotenv.Load(root + "/.env")
}
