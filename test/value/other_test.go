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
	require.Equal(t, []interface{}{"the_value"}, support.NewValue(support.NewCollection([]string{"the_value"})).Get("*").Raw())
}
