package test

import (
	"github.com/confetti-framework/support/str"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_upper_first_with_empty_string(t *testing.T) {
	result := str.UpperFirst("")
	require.Equal(t, "", result)
}

func Test_upper_first_with_multiple_words(t *testing.T) {
	result := str.UpperFirst("a horse is happy")
	require.Equal(t, "A horse is happy", result)
}

func Test_in_slice_with_no_parameter(t *testing.T) {
	require.False(t, str.InSlice("phone"))
}

func Test_in_slice_with_one_non_existing_string(t *testing.T) {
	require.False(t, str.InSlice("phone", "bag"))
}

func Test_in_slice_with_one_existing_string(t *testing.T) {
	require.True(t, str.InSlice("phone", "phone"))
}

func Test_in_slice_with_multiple_one_matched_parameters(t *testing.T) {
	require.True(t, str.InSlice("phone", "TV", "phone", "tabel"))
}

func Test_in_slice_with_integer(t *testing.T) {
	require.True(t, str.InSlice(1, 0, 1))
}

func Test_after(t *testing.T) {
	require.Equal(t, " my name", str.After("This is my name", "This is"))
	require.Equal(t, "nah", str.After("hannah", "han"))
	require.Equal(t, "nah", str.After("hannah", "n"))
	require.Equal(t, "nah", str.After("ééé hannah", "han"))
	require.Equal(t, "hannah", str.After("hannah", "xxxx"))
	require.Equal(t, "hannah", str.After("hannah", ""))
	require.Equal(t, "nah", str.After("han0nah", "0"))
}

func Test_afterLast(t *testing.T) {
	require.Equal(t, "tte", str.AfterLast("yvette", "yve"))
	require.Equal(t, "Controller", str.AfterLast("App\\Http\\Controllers\\Controller", "\\"))
	require.Equal(t, "e", str.AfterLast("yvette", "t"))
	require.Equal(t, "e", str.AfterLast("ééé yvette", "t"))
	require.Equal(t, "foo", str.AfterLast("----foo", "---"))
	require.Equal(t, "", str.AfterLast("yvette", "tte"))
	require.Equal(t, "yvette", str.AfterLast("yvette", "xxxx"))
	require.Equal(t, "yvette", str.AfterLast("yvette", ""))
	require.Equal(t, "te", str.AfterLast("yv0et0te", "0"))
}

func Test_before(t *testing.T) {
	require.Equal(t, "This is ", str.Before("This is my name", "my name"))
	require.Equal(t, "han", str.Before("hannah", "nah"))
	require.Equal(t, "ha", str.Before("hannah", "n"))
	require.Equal(t, "ééé ", str.Before("ééé hannah", "han"))
	require.Equal(t, "hannah", str.Before("hannah", "xxxx"))
	require.Equal(t, "hannah", str.Before("hannah", ""))
	require.Equal(t, "han", str.Before("han0nah", "0"))
}
