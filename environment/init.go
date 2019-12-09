package environment

import (
	"github.com/joho/godotenv"
	"github.com/lanvard/support/caller"
	"os"
	"path/filepath"
)

func init() {
	// load from project/vendor
	basePath := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(caller.CurrentDir()))))

	file := basePath + string(os.PathSeparator) + ".env"
	err := godotenv.Load(file)

	if err != nil {
		// load from GOROOT/vendor
		basePath := filepath.Dir(filepath.Dir(filepath.Dir(filepath.Dir(caller.CurrentDir()))))

		file := basePath + string(os.PathSeparator) + "lanvard/.env"
		err := godotenv.Load(file)

		if err != nil {
			println(err)
			panic("Error loading .env file in directory " + file)
		}
	}
}
