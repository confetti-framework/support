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
