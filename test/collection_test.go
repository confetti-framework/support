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

	value := values.Get("*").Collection()
	require.Len(
		t,
		value,
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

	name := values.Get("names.*").Collection()[0].Raw()
	require.Equal(
		t,
		"David",
		name,
	)
}
