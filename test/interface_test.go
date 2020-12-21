package test

import (
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

type testInterface interface{}
type testStruct struct{}

func Test_name_from_nil(t *testing.T) {
	name := support.Name((*testInterface)(nil))
	require.Equal(t, "test.testInterface", name)
}

func Test_name_from_struct(t *testing.T) {
	name := support.Name(testStruct{})
	require.Equal(t, "test.testStruct", name)
}

func Test_name_from_string(t *testing.T) {
	name := support.Name("InterfaceByString")
	require.Equal(t, "InterfaceByString", name)
}

func Test_type_from_interface(t *testing.T) {
	reflectType := support.Kind((*testInterface)(nil))
	require.Equal(t, reflect.Ptr, reflectType)
}

func Test_type_from_string(t *testing.T) {
	reflectType := support.Kind("string")
	require.Equal(t, reflect.String, reflectType)
}
