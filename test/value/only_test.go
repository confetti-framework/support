package value

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_map_only_from_empty_value(t *testing.T) {
	value := support.NewValue(nil)
	require.Equal(t, nil, value.Only().Raw())
}

func Test_map_only_from_value(t *testing.T) {
	value := support.NewValue(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.Equal(t, map[string]interface{}{"username": "apple_pear"}, value.Only("username").Raw())
}

func Test_collection_only_from_value(t *testing.T) {
	value := support.NewValue([]string{"salamander", "koala"})
	require.Equal(t, []interface{}{"salamander", "koala"}, value.Only("*").Raw())
}
