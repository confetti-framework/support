package support

import "strings"

type Errors []error

func (e Errors) Error() string {
	var errors []string
	for _, err := range e {
		errors = append(errors, err.Error())
	}

	return strings.Join(errors, " and ")
}
