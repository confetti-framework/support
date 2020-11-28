package value

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func Test_keys_empty(t *testing.T) {
	var given []string
	verboseKeys, settableKeys := support.GetSearchableKeys(given, support.NewValue(nil))
	require.Equal(t, []string{}, verboseKeys)
	require.Equal(t, []string{}, settableKeys)
}

func Test_keys_given_with_empty_value(t *testing.T) {
	given := []string{"name", "color"}
	verboseKeys, settableKeys := support.GetSearchableKeys(given, support.NewValue(nil))
	require.Equal(t, []string{"name", "color"}, verboseKeys)
	require.Equal(t, []string{"name", "color"}, settableKeys)
}

func Test_keys_given_with_map_one_value(t *testing.T) {
	given := []string{"name"}
	expect := []string{"name"}
	value := support.NewValue(support.NewMap(map[string]string{"name": "whale"}))
	verboseKeys, settableKeys := support.GetSearchableKeys(given, value)
	require.Equal(t, expect, verboseKeys)
	require.Equal(t, expect, settableKeys)
}

func Test_keys_given_with_map_two_values(t *testing.T) {
	given := []string{"name", "color"}
	expect := []string{"name", "color"}
	value := support.NewValue(support.NewMap(map[string]string{"name": "mule", "color": "black"}))
	verboseKeys, settableKeys := support.GetSearchableKeys(given, value)
	require.Equal(t, expect, verboseKeys)
	require.Equal(t, expect, settableKeys)
}

func Test_keys_asterisk_given_with_map(t *testing.T) {
	given := []string{"*"}
	expect := []string{"name", "color"}
	value := support.NewValue(support.NewMap(map[string]string{"name": "mule", "color": "black"}))
	verboseKeys, settableKeys := support.GetSearchableKeys(given, value)
	equalStrings(t, expect, verboseKeys)
	equalStrings(t, expect, settableKeys)
}

func Test_keys_map_with_second_map(t *testing.T) {
	given := []string{"*.name"}
	expect := []string{"big.name", "small.name"}
	input := map[string]map[string]string{"big": {"name": "whale"}, "small": {"name": "crab"}}

	value := support.NewValue(support.NewMap(input))
	verboseKeys, settableKeys := support.GetSearchableKeys(given, value)
	equalStrings(t, expect, verboseKeys)
	equalStrings(t, expect, settableKeys)
}

func Test_keys_map_with_2_layers(t *testing.T) {
	given := []string{"*.*.name"}
	expect := []string{"animal.big.name"}
	input := map[string]map[string]map[string]string{"animal": {"big": {"name": "whale"}}}

	value := support.NewValue(support.NewMap(input))
	verboseKeys, settableKeys := support.GetSearchableKeys(given, value)
	equalStrings(t, expect, verboseKeys)
	equalStrings(t, expect, settableKeys)
}

func Test_no_keys_with_map(t *testing.T) {
	given := []string{}
	expect := []string{}
	value := support.NewValue(support.NewMap(map[string]string{"name": "mule", "color": "black"}))
	verboseKeys, settableKeys := support.GetSearchableKeys(given, value)
	require.Equal(t, expect, verboseKeys)
	require.Equal(t, expect, settableKeys)
}

func Test_keys_with_collection(t *testing.T) {
	given := []string{"*"}
	value := support.NewValue(support.NewCollection([]string{"mule", "black"}))
	verboseKeys, settableKeys := support.GetSearchableKeys(given, value)
	require.Equal(t, []string{"0", "1"}, verboseKeys)
	require.Equal(t, []string{"*", "*"}, settableKeys)
}

func Test_keys_with_collection_and_map(t *testing.T) {
	given := []string{"*.*"}
	value := support.NewValue(support.NewCollection(support.NewMap(map[string]string{"big": "mule", "small": "black"})))
	verboseKeys, settableKeys := support.GetSearchableKeys(given, value)
	equalStrings(t, []string{"0.big", "0.small"}, verboseKeys)
	equalStrings(t, []string{"*.big", "*.small"}, settableKeys)
}

func equalStrings(t *testing.T, expect []string, result []string) {
	sort.Strings(expect)
	sort.Strings(result)
	require.Equal(t, expect, result)
}
