package support

type Key struct {
	// keys where collection key is a number
	searchable string
	// keys where collection key is an asterisks
	settable string
}

func (k Key) Searchable() string {
	return k.searchable
}

func (k Key) Settable() string {
	return k.settable
}

func (k Key) Wrap(searchablePrefix, settablePrefix string) Key {
	if k.searchable != "" {
		searchablePrefix = searchablePrefix + "."
	}
	k.searchable = searchablePrefix + k.searchable

	if k.settable != "" {
		settablePrefix = settablePrefix + "."
	}
	k.settable = settablePrefix + k.settable

	return k
}

func NewKey(key string) Key {
	return Key{searchable: key, settable: key}
}
