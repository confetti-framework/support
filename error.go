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

func Wrap(err error, message string) *Error {
	return &Error{err: errors.Wrap(err, message)}
}

func (e *Error) Error() string {
	return e.err.Error()
}

func (e *Error) StackTrace() errors.StackTrace {
	if errWithStack, ok := e.err.(interface{ StackTrace() errors.StackTrace }); ok {
		return errWithStack.StackTrace()
	}
	return nil
}

func (e *Error) Level(inputLevel level.Level) *Error {
	e.level = inputLevel
	return e
}

func (e *Error) GetLevel() level.Level {
	return e.level
}
