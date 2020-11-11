package test

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_raw_from_empty_string(t *testing.T) {
	assert.Equal(t, "", support.NewValue("").Raw())
}

func Test_raw_from_nil_string(t *testing.T) {
	assert.Equal(t, nil, support.NewValue(nil).Raw())
}

func Test_raw_from_string(t *testing.T) {
	assert.Equal(t, "flour", support.NewValue("flour").Raw())
}

func Test_raw_from_bool(t *testing.T) {
	assert.Equal(t, true, support.NewValue(true).Raw())
	assert.Equal(t, false, support.NewValue(false).Raw())
}

func Test_raw_from_number(t *testing.T) {
	assert.Equal(t, 100, support.NewValue(100).Raw())
	assert.Equal(t, -100, support.NewValue(-100).Raw())
}

func Test_raw_from_float(t *testing.T) {
	assert.Equal(t, 0.1, support.NewValue(0.1).Raw())
}

func Test_raw_from_collection_with_one_string(t *testing.T) {
	assert.Equal(t, []interface{}{"door"}, support.NewCollection("door").Raw())
}

func Test_raw_from_collection_with_tho_strings(t *testing.T) {
	assert.Equal(t, []interface{}{"foo", "bar"}, support.NewCollection("foo", "bar").Raw())
}

func Test_raw_from_collection_with_tho_numbers(t *testing.T) {
	assert.Equal(t, []interface{}{12, 14}, support.NewCollection(12, 14).Raw())
}

func Test_raw_from_collection_with_tho_float(t *testing.T) {
	assert.Equal(t, []interface{}{1.5, 0.4}, support.NewCollection(1.5, 0.4).Raw())
}

func Test_raw_from_value_with_collection(t *testing.T) {
	actual := support.NewValue(support.NewCollection("door")).Raw()
	assert.Equal(t, []interface{}{"door"}, actual)
}

func Test_raw_from_map_with_strings(t *testing.T) {
	actual := support.NewMap(map[string]string{
		"chair": "blue",
		"table": "green",
	}).Raw()

	assert.Equal(t, map[string]interface{}{"chair": "blue", "table": "green"}, actual)
}

func Test_raw_from_value_with_collection_and_map(t *testing.T) {
	actual := support.NewValue(
		support.NewCollection(
			support.NewMap(map[string]string{"key": "door"}),
		),
	).Raw()

	assert.Equal(t, []interface{}{map[string]interface{}{"key": "door"}}, actual)
}

func Test_raw_from_value_with_error(t *testing.T) {
	raw, err := support.NewValueE(100, "this is an error").RawE()
	assert.Equal(t, 100, raw)
	assert.EqualError(t, err, "this is an error")
}

func Test_raw_from_value_without_error(t *testing.T) {
	raw, err := support.NewValueE(100, nil).RawE()
	assert.Equal(t, 100, raw)
	assert.NoError(t, err)
}

func Test_raw_from_value_and_collection_with_multiple_errors(t *testing.T) {
	raw, errs := support.NewValue(
		support.NewCollection(
			support.NewValueE(100, "this the first error"),
			support.NewValueE(100, "this is the second error"),
		),
	).RawE()

	assert.Equal(t, []interface{}{100, 100}, raw)
	assert.EqualError(t, errs, "this is the second error")
}

func Test_raw_from_value_and_map_with_multiple_errors(t *testing.T) {
	raw, errs := support.NewValue(
		support.NewMap(map[string]interface{}{
			"key1": support.NewValueE(100, "this the first error"),
			"key2": support.NewValueE(150, "this is the second error"),
		}),
	).RawE()

	assert.Equal(t, map[string]interface{}{"key1": 100, "key2": 150}, raw)
	assert.EqualError(t, errs, "this is the second error")
}

func Test_raw_value_panic(t *testing.T) {
	value := support.NewValueE(nil, "an error")

	action := func() {
		value.Raw()
	}

	require.PanicsWithError(t, "an error", action)
}

func Test_raw_value_with_value_panic(t *testing.T) {
	value := support.NewValue(support.NewValueE(nil, "an error"))

	action := func() {
		value.Raw()
	}

	require.PanicsWithError(t, "an error", action)
}

func Test_raw_collection_panic(t *testing.T) {
	value := support.NewCollection(support.NewValueE(nil, "an error"))

	action := func() {
		value.Raw()
	}

	require.PanicsWithError(t, "an error", action)
}

func Test_raw_map_panic(t *testing.T) {
	value := support.NewMap(map[string]interface{}{
		"key": support.NewValueE(nil, "an error"),
	})

	action := func() {
		value.Raw()
	}

	require.PanicsWithError(t, "an error", action)
}
