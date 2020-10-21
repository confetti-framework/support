package support

import (
	"github.com/lanvard/syslog/level"
	"github.com/pkg/errors"
)

type Error struct {
	err   error
	level level.Level
}

func NewError(message string) *Error {
	return &Error{err: errors.New(message)}
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) Level(inputLevel level.Level) *Error {
	e.level = inputLevel
	return e
}

func (e *Error) GetLevel() level.Level {
	return e.level
}
