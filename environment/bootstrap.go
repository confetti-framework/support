package environment

import (
	"github.com/joho/godotenv"
	"github.com/lanvard/contract/inter"
	"os"
	"strings"
)

func Bootstrap(basePath inter.BasePath) error  {
	var file string

	if strings.HasSuffix(os.Args[0], ".test") {
		file = basePath.EnvironmentTestingFile()
	} else {
		file = basePath.EnvironmentFile()
	}

	return godotenv.Load(file)
}
