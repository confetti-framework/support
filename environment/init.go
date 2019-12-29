package environment

import (
	"github.com/joho/godotenv"
	"lanvard/config/entity"
	"os"
	"strings"
)

func init() {
	var file string

	if strings.HasSuffix(os.Args[0], ".test") {
		file = entity.NewBasePath().EnvironmentFile()
	} else {
		file = entity.NewBasePath().EnvironmentTestingFile()
	}

	err := godotenv.Load(file)

	if err != nil {
		println(err)
		panic("Error loading .env file in directory " + file)
	}
}
