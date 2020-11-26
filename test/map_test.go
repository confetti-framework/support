package test

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_get_all_from_map(t *testing.T) {
	values := support.NewMap(map[string][]string{
		"language": {"Go"},
		"names":    {"David", "Jona"},
	})

	value := values.Get("*")
	require.Len(
		t,
		value.Collection(),
		3,
	)
}

func Test_map_only_when_all_keys_are_present(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.Equal(t, data, data.Only("username", "password"))
}

func Test_map_only_when_less_keys_than_present(t *testing.T) {
	require.Equal(
		t,
		support.NewMap(map[string]string{"username": "apple_pear"}),
		support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"}).Only("username"),
	)
}

func Test_map_only_when_more_keys_than_present(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.Equal(t, data, data.Only("username", "password", "age"))
}

func Test_map_only_with_asterisk_and_nested_key(t *testing.T) {
	data := support.NewMap(map[string]map[string]string{
		"piet_niet": {
			"username": "piet",
			"password": "afd23432a12",
		},
		"jan_kan": {
			"username": "jan",
			"password": "34a@#dQd",
		},
	})

	require.Equal(t, map[string]interface{}{
		"piet_niet": map[string]interface{}{
			"username": "piet",
		},
		"jan_kan": map[string]interface{}{
			"username": "jan",
		}}, data.Only("*.username").Raw())
}

func Test_map_except_when_no_keys_are_present(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.Equal(t, data, data.Except())
}

func Test_map_except_when_less_keys_than_present(t *testing.T) {
	require.Equal(t,
		support.NewMap(map[string]string{"username": "apple_pear"}),
		support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"}).Except("password"),
	)
}

func Test_map_except_when_more_keys_than_present(t *testing.T) {
	require.Equal(t,
		support.NewMap(map[string]string{"username": "apple_pear"}),
		support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"}).Except("password", "age"),
	)
}

func Test_no_reference_to_old_struct(t *testing.T) {
	oldStruct := support.NewMap(map[string]string{"username": "apple_pear"})
	_ = oldStruct.Except("username")
	require.Equal(t, oldStruct, support.NewMap(map[string]string{"username": "apple_pear"}))
}

func Test_map_has_with_no_key(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear"})
	require.True(t, data.Has())
}

func Test_map_has_with_one_key(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear"})
	require.True(t, data.Has("username"))
	require.False(t, data.Has("age"))
}

func Test_map_has_with_multiple_keys(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.True(t, data.Has("username", "password"))
	require.False(t, data.Has("username", "age"))
}

func Test_map_has_with_nil_value(t *testing.T) {
	user := map[string]support.Value{"user": support.NewValue(nil)}
	data := support.NewMap(user)
	require.True(t, data.Has("user"))
}

func Test_map_has_any_with_no_key(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.True(t, data.HasAny())
}

func Test_map_has_any_with_one_non_present_key(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.False(t, data.HasAny("age"))
}

func Test_map_missing_with_no_key(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear"})
	require.False(t, data.Missing(""))
}
func Test_map_missing_with_one_key(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.False(t, data.Missing("username"))
}

func Test_map_missing_one_key_missing(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.True(t, data.Missing("age", "password"))
}

func Test_map_filled_with_no_key(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear"})
	require.True(t, data.Filled())
}

func Test_map_filled_with_one_key(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.True(t, data.Filled("username"))
}

func Test_map_filled_with_multiple_key_filled(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	require.True(t, data.Filled("username", "password"))
}

func Test_map_filled_with_one_not_filled_but_present(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear", "password": ""})
	require.False(t, data.Filled("username", "password"))
}

func Test_map_filled_with_one_not_present(t *testing.T) {
	data := support.NewMap(map[string]string{"username": "apple_pear"})
	require.False(t, data.Filled("username", "password"))
}

func Test_map_set_value(t *testing.T) {
	data := support.NewMap(map[string]string{})
	data.SetE("username", support.NewValue("apple_pear"))
	require.Equal(t, "apple_pear", data.Get("username").String())
}

func Test_map_set_string(t *testing.T) {
	data := support.NewMap(map[string]string{})
	data.SetE("username", "apple_pear")
	require.Equal(t, "apple_pear", data.Get("username").String())
}

func Test_map_set_struct(t *testing.T) {
	data := support.NewMap(map[string]string{})
	data.SetE("user", mockUser{})
	require.Equal(t, mockUser{}, data.Get("user").Raw())
}

func Test_map_set_by_dot_notation(t *testing.T) {
	data := support.NewMap(map[string]string{})
	data.SetE("user.name", "Rob")
	require.Equal(t, "Rob", data.Get("user.name").String())
}

func Test_map_set_by_dot_notation_with_existing_data(t *testing.T) {
	data := support.NewMap(map[string]interface{}{"user": map[string]string{"street": "Frozen street"}})
	data.SetE("user.name", "Rob")
	require.Equal(t, "Rob", data.Get("user.name").String())
	require.Equal(t, "Frozen street", data.Get("user.street").String())
	require.Equal(t, map[string]interface{}{"street": "Frozen street", "name": "Rob"}, data.Get("user").Map().Raw())
}

func Test_set_collection_on_map(t *testing.T) {
	var err error
	data := support.NewValue(nil)
	data, err = data.SetE("names.*", "Jaap")
	require.Nil(t, err)
	require.Equal(t, map[string]interface{}{"names": []interface{}{"Jaap"}}, data.Raw())
}

type mockUser struct{}
