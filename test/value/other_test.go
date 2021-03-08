package value

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockEmptyStruct struct{}
type mockStruct struct {
	Field string
}

var mockFunc = func() {}

func Test_new_invalid_value(t *testing.T) {
	require.Panics(t, func() {
		support.NewValue(map[interface{}]string{
			mockFunc: "val",
		})
	})
}

func Test_get_collection_from_value_with_asterisks(t *testing.T) {
	value := support.NewValue(support.NewCollection([]string{"the_value"}))
	require.Equal(t, []interface{}{"the_value"}, value.Get("*").Raw())
}

func Test_get_map_from_value_with_asterisks(t *testing.T) {
	value := support.NewValue(support.NewMap(map[string]interface{}{"key": "the_value"}))
	require.Equal(t, map[string]interface{}{"key": "the_value"}, value.Get("*").Raw())
}

func Test_get_from_invalid_value_with_asterisks(t *testing.T) {
	value := support.NewValue("non collection/value")
	require.Panics(t, func() {
		value.Get("*").Raw()
	})
}

func Test_get_from_empty_struct(t *testing.T) {
	value := support.NewValue(mockEmptyStruct{})
	require.Panics(t, func() {
		value.Get("field").Raw()
	})
}

func Test_get_from_struct(t *testing.T) {
	value := support.NewValue(mockStruct{"fieldvalue"})
	require.Equal(t, "fieldvalue", value.Get("Field").Raw())
}

func Test_get_from_int(t *testing.T) {
	value := support.NewValue(12)
	require.Panics(t, func() {
		value.Get("field").Raw()
	})
}
