package value

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"testing"
)

var mockConst = func() {}

func Test_new_invalid_value(t *testing.T) {
	require.Panics(t, func() {
		support.NewValue(map[interface{}]string{
			mockConst: "val",
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
