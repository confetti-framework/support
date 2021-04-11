package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAfter(t *testing.T) {
	// TODO: What if nothing is found?
	assert.Equal(t, "", After("", ""))
	assert.Equal(t, "", After("hannah", ""))
	assert.Equal(t, "", After("", "han"))
	assert.Equal(t, "nah", After("hannah", "han"))
	assert.Equal(t, "nah", After("hannah", "n"))
	assert.Equal(t, "nah", After("eee hannah", "han"))
	assert.Equal(t, "nah", After("ééé hannah", "han"))
	assert.Equal(t, "hannah", After("hannah", "xxxx"))
	assert.Equal(t, "nah", After("han0nah", "0"))
	assert.Equal(t, "nah", After("han2nah", "2"))
}

func TestAfterLast(t *testing.T) {
	// TODO: What if nothing is found?
	assert.Equal(t, "", After("", ""))
	assert.Equal(t, "", After("hannah", ""))
	assert.Equal(t, "", After("", "han"))
	assert.Equal(t,"tte", AfterLast("yvette", "yve"))
	assert.Equal(t,"e", AfterLast("yvette", "t"))
	assert.Equal(t,"e", AfterLast("ééé yvette", "t"))
	assert.Equal(t,"", AfterLast("yvette", "tte"))
	assert.Equal(t,"yvette", AfterLast("yvette", "xxxx"))
	assert.Equal(t,"te", AfterLast("yv0et0te", "0"))
	assert.Equal(t,"te", AfterLast("yv0et0te", "0"))
	assert.Equal(t,"te", AfterLast("yv2et2te", "2"))
	assert.Equal(t,"foo", AfterLast("----foo", "---"))
}

func TestBefore(t *testing.T) {
	assert.Equal(t, "han", Before("hannah", "nah"))
	assert.Equal(t, "ha", Before("hannah", "n"))
	assert.Equal(t, "ééé ", Before("ééé hannah", "han"))
	assert.Equal(t, "hannah", Before("hannah", "xxxx"))
	assert.Equal(t, "han", Before("han0nah", "0"))
	assert.Equal(t, "han", Before("han0nah", "0"))
	assert.Equal(t, "han", Before("han2nah", "2"))
}

func TestBeforeLast(t *testing.T) {
	assert.Equal(t,"yve", BeforeLast("yvette", "tte"))
	assert.Equal(t,"yvet", BeforeLast("yvette", "t"))
	assert.Equal(t,"ééé ", BeforeLast("ééé yvette", "yve"))
	assert.Equal(t,"", BeforeLast("yvette", "yve"))
	assert.Equal(t,"yvette", BeforeLast("yvette", "xxxx"))
	assert.Equal(t,"yvette", BeforeLast("yvette", ""))
	assert.Equal(t,"yv0et", BeforeLast("yv0et0te", "0"))
	assert.Equal(t,"yv0et", BeforeLast("yv0et0te", "0"))
	assert.Equal(t,"yv2et", BeforeLast("yv2et2te", "2"))
}

func TestBetween(t *testing.T) {
	assert.Equal(t,"abc", Between("abc", "", "c"))
	assert.Equal(t,"abc", Between("abc", "a", ""))
	assert.Equal(t,"abc", Between("abc", "", ""))
	assert.Equal(t,"b", Between("abc", "a", "c"))
	assert.Equal(t,"b", Between("dddabc", "a", "c"))
	assert.Equal(t,"b", Between("abcddd", "a", "c"))
	assert.Equal(t,"b", Between("dddabcddd", "a", "c"))
	assert.Equal(t,"nn", Between("hannah", "ha", "ah"))
	assert.Equal(t,"a]ab[b", Between("[a]ab[b]", "[", "]"))
	assert.Equal(t,"foo", Between("foofoobar", "foo", "bar"))
	assert.Equal(t,"bar", Between("foobarbar", "foo", "bar"))
}


