package str

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAfter(t *testing.T) {

	assert.Equal(t, "nah", After("hannah", "han"))
	assert.Equal(t, "nah", After("hannah", "n"))
	assert.Equal(t, "nah", After("eee hannah", "han"))
	assert.Equal(t, "nah", After("ééé hannah", "han"))
	assert.Equal(t, "hannah", After("hannah", "xxxx"))
	assert.Equal(t, "hannah", After("hannah", ""))
	assert.Equal(t, "nah", After("han0nah", "0"))
	assert.Equal(t, "nah", After("han2nah", "2"))
}

func TestAfterLast(t *testing.T) {
	assert.Equal(t,"tte", AfterLast("yvette", "yve"))
	assert.Equal(t,"e", AfterLast("yvette", "t"))
	assert.Equal(t,"e", AfterLast("ééé yvette", "t"))
	assert.Equal(t,"", AfterLast("yvette", "tte"))
	assert.Equal(t,"yvette", AfterLast("yvette", "xxxx"))
	assert.Equal(t,"yvette", AfterLast("yvette", ""))
	assert.Equal(t,"te", AfterLast("yv0et0te", "0"))
	assert.Equal(t,"te", AfterLast("yv0et0te", "0"))
	assert.Equal(t,"te", AfterLast("yv2et2te", "2"))
	assert.Equal(t,"foo", AfterLast("----foo", "---"))
}
