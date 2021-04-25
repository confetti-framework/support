package value

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

var theErr = errors.New("the error")

func Test_error_to_string(t *testing.T) {
	value, err := support.NewValue(nil, theErr).StringE()
	assert.Equal(t, "", value)
	assert.Error(t, err)
}