package test

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllFromCollection(t *testing.T) {
	values := support.NewCollection([]string{
		"Go",
		"David",
		"Sammy",
	})

	value := values.Get("*")
	assert.Len(
		t,
		value.Collection(),
		3,
	)
}

func TestGetCollectionByKey(t *testing.T) {
	values := support.NewMap(map[string][]string{
		"language": {"Go"},
		"names":    {"David", "Jona"},
	})

	languages := values.Get("language.*")
	assert.Equal(t, "Go", languages.Collection().First().String())

	names := values.Get("names.*")
	assert.Equal(
		t,
		support.NewCollection([]string{"David", "Jona"}),
		names.Collection(),
	)
}
