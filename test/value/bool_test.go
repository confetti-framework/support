package value

import (
	"github.com/lanvard/errors"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolTrueFromTrue(t *testing.T) {
	assert.True(t, support.NewValue(true).Bool())
}

func TestBoolTrueFromFalse(t *testing.T) {
	assert.False(t, support.NewValue(false).Bool())
}

func TestBoolTrueFromIntOne(t *testing.T) {
	assert.True(t, support.NewValue(1).Bool())
}

func TestBoolTrueFromIntTwo(t *testing.T) {
	assert.False(t, support.NewValue(2).Bool())
}

func TestBoolTrueFromIntZero(t *testing.T) {
	assert.False(t, support.NewValue(0).Bool())
}

func TestBoolTrueFromStringOne(t *testing.T) {
	assert.True(t, support.NewValue("1").Bool())
}

func TestBoolTrueFromStringTrue(t *testing.T) {
	assert.True(t, support.NewValue("true").Bool())
}

func TestBoolTrueFromStringOn(t *testing.T) {
	assert.True(t, support.NewValue("on").Bool())
}

func TestBoolTrueFromStringYes(t *testing.T) {
	assert.True(t, support.NewValue("yes").Bool())
}

func TestBoolTrueWithErrorShouldPanic(t *testing.T) {
	assert.PanicsWithError(t, "the error", func() {
		support.NewValueE("yes", errors.New("the error")).Bool()
	})
}

func TestBoolWithFalseAndError(t *testing.T) {
	value, err := support.NewValueE(false, errors.New("the error")).BoolE()

	assert.False(t, value)
	assert.EqualError(t, err, "the error")
}

func TestBoolWithTrueAndError(t *testing.T) {
	value, err := support.NewValueE(true, errors.New("the error")).BoolE()

	assert.True(t, value)
	assert.EqualError(t, err, "the error")
}
