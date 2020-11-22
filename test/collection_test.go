package test

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_get_all_from_collection(t *testing.T) {
	values := support.NewCollection([]string{
		"Go",
		"David",
		"Sammy",
	})

	value := values.Get("*").Collection()
	require.Len(
		t,
		value,
		3,
	)
}

func Test_get_collection_by_key(t *testing.T) {
	values := support.NewMap(map[string][]string{
		"language": {"Go"},
		"names":    {"David", "Jona"},
	})

	languages := values.Get("language.*")
	require.Equal(t, "Go", languages.Collection().First().String())

	name := values.Get("names.*").Collection()[0].Raw()
	require.Equal(
		t,
		"David",
		name,
	)
}

func Test_get_collection_by_unknown_key(t *testing.T) {
	values := support.NewValue([]string{"house", "door"})

	result, err := values.GetE("2")
	require.EqualError(t, err, "key '2': can not found value")
	require.Equal(t, support.NewValue(emptyInterface), result)
}

func Test_get_collection_by_known_key(t *testing.T) {
	values := support.NewValue([]string{"house", "door"})

	result, err := values.GetE("1")
	require.Nil(t, err)
	require.Equal(t, "door", result.String())
}

func Test_collection_get_by_string(t *testing.T) {
	data := support.NewCollection([]string{})
	result, err := data.GetE("username")
	require.EqualError(t, err, "'username' can only be a number or *")
	require.Equal(t, support.NewValue(emptyInterface), result)
}

func Test_collection_push_value(t *testing.T) {
	data := support.NewCollection([]string{})
	data = data.Push(support.NewValue("apple_pear"))
	require.Equal(t, "apple_pear", data.Get("0").String())
}

func Test_collection_set_by_invalid_key(t *testing.T) {
	data := support.NewCollection([]string{})
	data, err := data.SetE("invalid_key", support.NewValue("apple_pear"))
	require.EqualError(t, err, "key 'invalid_key' can only begin with an asterisk")
	require.Equal(t, support.NewCollection([]string{}), data)
}

func Test_collection_set_asterisk(t *testing.T) {
	data := support.NewCollection([]string{})
	data, err := data.SetE("*", support.NewValue("apple_pear"))
	require.Nil(t, err)
	require.Equal(t, []interface{}{"apple_pear"}, data.Raw())
}

func Test_collection_set_nested_collection(t *testing.T) {
	data := support.NewCollection([]string{})
	data, err := data.SetE("*.*", support.NewValue("apple_pear"))
	require.Nil(t, err)
	require.Equal(t, []interface{}{[]interface{}{"apple_pear"}}, data.Raw())
}

func Test_collection_set_nested_collection_with_existing_data(t *testing.T) {
	data := support.NewCollection([]string{"berry"})
	data, err := data.SetE("*.*", support.NewValue("apple_pear"))
	require.Nil(t, err)
	require.Equal(t, []interface{}{"berry", []interface{}{"apple_pear"}}, data.Raw())
}

func Test_collection_set_collection(t *testing.T) {
	data := support.NewCollection()
	data, err := data.SetE("*", support.NewCollection(support.NewValue("apple_pear")))
	require.Nil(t, err)
	require.Equal(t, []interface{}{[]interface{}{"apple_pear"}}, data.Raw())
}

var emptyInterface interface{}
