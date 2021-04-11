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

func Test_After(t *testing.T) {
	// TODO: What if nothing is found?
	require.Equal(t, "", str.After("", ""))
	require.Equal(t, "", str.After("", "han"))
	require.Equal(t, "hannah", str.After("hannah", ""))
	require.Equal(t, "nah", str.After("hannah", "han"))
	require.Equal(t, "nah", str.After("hannah", "n"))
	require.Equal(t, "nah", str.After("eee hannah", "han"))
	require.Equal(t, "nah", str.After("ééé hannah", "han"))
	require.Equal(t, "hannah", str.After("hannah", "xxxx"))
	require.Equal(t, "nah", str.After("han0nah", "0"))
	require.Equal(t, "nah", str.After("han2nah", "2"))
}

func Test_AfterLast(t *testing.T) {
	// TODO: What if nothing is found?
	require.Equal(t, "", str.After("", ""))
	require.Equal(t, "", str.After("", "han"))
	require.Equal(t, "hannah", str.After("hannah", ""))
	require.Equal(t, "tte", str.AfterLast("yvette", "yve"))
	require.Equal(t, "e", str.AfterLast("yvette", "t"))
	require.Equal(t, "e", str.AfterLast("ééé yvette", "t"))
	require.Equal(t, "", str.AfterLast("yvette", "tte"))
	require.Equal(t, "yvette", str.AfterLast("yvette", "xxxx"))
	require.Equal(t, "te", str.AfterLast("yv0et0te", "0"))
	require.Equal(t, "te", str.AfterLast("yv0et0te", "0"))
	require.Equal(t, "te", str.AfterLast("yv2et2te", "2"))
	require.Equal(t, "foo", str.AfterLast("----foo", "---"))
}

func Test_Before(t *testing.T) {
	require.Equal(t, "han", str.Before("hannah", "nah"))
	require.Equal(t, "ha", str.Before("hannah", "n"))
	require.Equal(t, "ééé ", str.Before("ééé hannah", "han"))
	require.Equal(t, "hannah", str.Before("hannah", "xxxx"))
	require.Equal(t, "han", str.Before("han0nah", "0"))
	require.Equal(t, "han", str.Before("han0nah", "0"))
	require.Equal(t, "han", str.Before("han2nah", "2"))
}

func Test_BeforeLast(t *testing.T) {
	require.Equal(t, "yve", str.BeforeLast("yvette", "tte"))
	require.Equal(t, "yvet", str.BeforeLast("yvette", "t"))
	require.Equal(t, "ééé ", str.BeforeLast("ééé yvette", "yve"))
	require.Equal(t, "", str.BeforeLast("yvette", "yve"))
	require.Equal(t, "yvette", str.BeforeLast("yvette", "xxxx"))
	require.Equal(t, "yvette", str.BeforeLast("yvette", ""))
	require.Equal(t, "yv0et", str.BeforeLast("yv0et0te", "0"))
	require.Equal(t, "yv0et", str.BeforeLast("yv0et0te", "0"))
	require.Equal(t, "yv2et", str.BeforeLast("yv2et2te", "2"))
}

func Test_Between(t *testing.T) {
	require.Equal(t, "abc", str.Between("abc", "", "c"))
	require.Equal(t, "abc", str.Between("abc", "a", ""))
	require.Equal(t, "abc", str.Between("abc", "", ""))
	require.Equal(t, "b", str.Between("abc", "a", "c"))
	require.Equal(t, "b", str.Between("dddabc", "a", "c"))
	require.Equal(t, "b", str.Between("abcddd", "a", "c"))
	require.Equal(t, "b", str.Between("dddabcddd", "a", "c"))
	require.Equal(t, "nn", str.Between("hannah", "ha", "ah"))
	require.Equal(t, "a]ab[b", str.Between("[a]ab[b]", "[", "]"))
	require.Equal(t, "foo", str.Between("foofoobar", "foo", "bar"))
	require.Equal(t, "bar", str.Between("foobarbar", "foo", "bar"))
}

func Test_Contains(t *testing.T) {
	require.True(t, str.Contains("taylor", "ylo"))
	require.True(t, str.Contains("taylor", "taylor"))

	require.False(t, str.Contains("taylor", "xxx"))
	require.False(t, str.Contains("taylor", ""))
	require.False(t, str.Contains("", ""))
}
