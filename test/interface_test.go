package test

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type testInterface interface{}
type testStruct struct{}

func Test_name_from_nil(t *testing.T) {
	name := support.Name((*testInterface)(nil))
	assert.Equal(t, "test.testInterface", name)
}

func Test_name_from_struct(t *testing.T) {
	name := support.Name(testStruct{})
	assert.Equal(t, "test.testStruct", name)
}

func Test_name_from_string(t *testing.T) {
	name := support.Name("InterfaceByString")
	assert.Equal(t, "InterfaceByString", name)
}

func Test_type_from_interface(t *testing.T) {
	reflectType := support.Type((*testInterface)(nil))
	assert.Equal(t, reflect.Ptr, reflectType)
}

func Test_type_from_string(t *testing.T) {
	reflectType := support.Type("string")
	assert.Equal(t, reflect.String, reflectType)
}
