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

	value := values.Get("*")
	require.Len(
		t,
		value.Collection(),
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

	names := values.Get("names.*")
	require.Equal(
		t,
		support.NewCollection([]string{"David", "Jona"}),
		names.Collection(),
	)
}
