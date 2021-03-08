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
