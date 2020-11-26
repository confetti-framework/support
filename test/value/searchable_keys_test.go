package value

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func Test_keys_empty(t *testing.T) {
	var given []string
	result := support.GetSearchableKeys(given, support.NewValue(nil))
	require.Equal(t, []string{}, result)
}

func Test_keys_given_with_empty_value(t *testing.T) {
	given := []string{"name", "color"}
	result := support.GetSearchableKeys(given, support.NewValue(nil))
	require.Equal(t, []string{"name", "color"}, result)
}

func Test_keys_given_with_map_one_value(t *testing.T) {
	given := []string{"name"}
	expect := []string{"name"}
	value := support.NewValue(support.NewMap(map[string]string{"name": "whale"}))
	result := support.GetSearchableKeys(given, value)
	require.Equal(t, expect, result)
}

func Test_keys_given_with_map_two_values(t *testing.T) {
	given := []string{"name", "color"}
	expect := []string{"name", "color"}
	value := support.NewValue(support.NewMap(map[string]string{"name": "mule", "color": "black"}))
	result := support.GetSearchableKeys(given, value)
	require.Equal(t, expect, result)
}

func Test_keys_asterisk_given_with_map(t *testing.T) {
	given := []string{"*"}
	expect := []string{"name", "color"}
	value := support.NewValue(support.NewMap(map[string]string{"name": "mule", "color": "black"}))
	result := support.GetSearchableKeys(given, value)
	equalStrings(t, expect, result)
}

func Test_keys_map_with_second_map(t *testing.T) {
	given := []string{"*.name"}
	expect := []string{"big.name", "small.name"}
	input := map[string]map[string]string{"big": {"name": "whale"}, "small": {"name": "crab"}}

	value := support.NewValue(support.NewMap(input))
	result := support.GetSearchableKeys(given, value)
	equalStrings(t, expect, result)
}

func Test_keys_map_with_2_layers(t *testing.T) {
	given := []string{"*.*.name"}
	expect := []string{"animal.big.name"}
	input := map[string]map[string]map[string]string{"animal": {"big": {"name": "whale"}}}

	value := support.NewValue(support.NewMap(input))
	result := support.GetSearchableKeys(given, value)
	equalStrings(t, expect, result)
}

func Test_no_keys_with_map(t *testing.T) {
	given := []string{}
	expect := []string{}
	value := support.NewValue(support.NewMap(map[string]string{"name": "mule", "color": "black"}))
	result := support.GetSearchableKeys(given, value)
	require.Equal(t, expect, result)
}

func Test_keys_with_collection(t *testing.T) {
	given := []string{"*"}
	expect := []string{"*"}
	value := support.NewValue(support.NewCollection([]string{"mule", "black"}))
	result := support.GetSearchableKeys(given, value)
	require.Equal(t, expect, result)
}

func equalStrings(t *testing.T, expect []string, result []string) {
	sort.Strings(expect)
	sort.Strings(result)
	require.Equal(t, expect, result)
}
