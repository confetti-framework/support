module github.com/lanvard/support

go 1.15

require (
	github.com/joho/godotenv v1.3.0
	github.com/lanvard/errors v0.9.1
	github.com/spf13/cast v1.3.1
	github.com/stretchr/testify v1.6.1
	golang.org/x/text v0.3.2
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace (
	github.com/lanvard/errors v0.9.1 => ../errors
	github.com/lanvard/syslog v0.0.0-20201006215111-98d4d91dbaa8 => ../syslog
)
