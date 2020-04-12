package environment

import (
	"github.com/joho/godotenv"
	"github.com/lanvard/contract/inter"
	"os"
	"strings"
)

func BootstrapEnvironment(basePath inter.BasePath)  {
	var file string

	if strings.HasSuffix(os.Args[0], ".test") {
		file = basePath.EnvironmentTestingFile()
	} else {
		file = basePath.EnvironmentFile()
	}

	err := godotenv.Load(file)

	if err != nil {
		println(err)
		panic("Error loading .env file in directory " + file)
	}
}
