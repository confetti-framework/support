package test

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRawFromEmptyString(t *testing.T) {
	assert.Equal(t, "", support.NewValue("").Raw())
}

func TestRawFromNilString(t *testing.T) {
	assert.Equal(t, nil, support.NewValue(nil).Raw())
}

func TestRawFromString(t *testing.T) {
	assert.Equal(t, "flour", support.NewValue("flour").Raw())
}

func TestRawFromBool(t *testing.T) {
	assert.Equal(t, true, support.NewValue(true).Raw())
	assert.Equal(t, false, support.NewValue(false).Raw())
}

func TestRawFromNumber(t *testing.T) {
	assert.Equal(t, 100, support.NewValue(100).Raw())
	assert.Equal(t, -100, support.NewValue(-100).Raw())
}

func TestRawFromFloat(t *testing.T) {
	assert.Equal(t, 0.1, support.NewValue(0.1).Raw())
}

func TestRawFromCollectionWithOneString(t *testing.T) {
	assert.Equal(t, []interface{}{"door"}, support.NewCollection("door").Raw())
}

func TestRawFromCollectionWithThoStrings(t *testing.T) {
	assert.Equal(t, []interface{}{"foo", "bar"}, support.NewCollection("foo", "bar").Raw())
}

func TestRawFromCollectionWithThoNumbers(t *testing.T) {
	assert.Equal(t, []interface{}{12, 14}, support.NewCollection(12, 14).Raw())
}

func TestRawFromCollectionWithThoFloat(t *testing.T) {
	assert.Equal(t, []interface{}{1.5, 0.4}, support.NewCollection(1.5, 0.4).Raw())
}

func TestRawFromValueWithCollection(t *testing.T) {
	actual := support.NewValue(support.NewCollection("door")).Raw()
	assert.Equal(t, []interface{}{"door"}, actual)
}

func TestRawFromMapWithStrings(t *testing.T) {
	actual := support.NewMap(map[string]string{
		"chair": "blue",
		"table": "green",
	}).Raw()

	assert.Equal(t, map[string]interface{}{"chair": "blue", "table": "green"}, actual)
}

func TestRawFromValueWithCollectionAndMap(t *testing.T) {
	actual := support.NewValue(
		support.NewCollection(
			support.NewMap(map[string]string{"key": "door"}),
		),
	).Raw()

	assert.Equal(t, []interface{}{map[string]interface{}{"key": "door"}}, actual)
}

func TestRawFromValueWithError(t *testing.T) {
	raw, err := support.NewValueE(100, "this is an error").RawE()
	assert.Equal(t, 100, raw)
	assert.EqualError(t, err, "this is an error")
}

func TestRawFromValueWithoutError(t *testing.T) {
	raw, err := support.NewValueE(100, nil).RawE()
	assert.Equal(t, 100, raw)
	assert.Nil(t, err)
}

func TestRawFromValueAndCollectionWithMultipleErrors(t *testing.T) {
	raw, errs := support.NewValue(
		support.NewCollection(
			support.NewValueE(100, "this the first error"),
			support.NewValueE(100, "this is the second error"),
		),
	).RawE()

	assert.Equal(t, []interface{}{100, 100}, raw)
	assert.EqualError(t, errs, "this is the second error")
}

func TestRawFromValueAndMapWithMultipleErrors(t *testing.T) {
	raw, errs := support.NewValue(
		support.NewMap(map[string]interface{}{
			"key1": support.NewValueE(100, "this the first error"),
			"key2": support.NewValueE(150, "this is the second error"),
		}),
	).RawE()

	assert.Equal(t, map[string]interface{}{"key1": 100, "key2": 150}, raw)
	assert.EqualError(t, errs, "this is the second error")
}

func TestRawValuePanic(t *testing.T) {
	value := support.NewValueE(nil, "an error")

	action := func() {
		value.Raw()
	}

	require.PanicsWithError(t, "an error", action)
}

func TestRawValueWithValuePanic(t *testing.T) {
	value := support.NewValue(support.NewValueE(nil, "an error"))

	action := func() {
		value.Raw()
	}

	require.PanicsWithError(t, "an error", action)
}

func TestRawCollectionPanic(t *testing.T) {
	value := support.NewCollection(support.NewValueE(nil, "an error"))

	action := func() {
		value.Raw()
	}

	require.PanicsWithError(t, "an error", action)
}

func TestRawMapPanic(t *testing.T) {
	value := support.NewMap(map[string]interface{}{
		"key": support.NewValueE(nil, "an error"),
	})

	action := func() {
		value.Raw()
	}

	require.PanicsWithError(t, "an error", action)
}
