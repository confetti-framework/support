package support

import (
	"github.com/lanvard/contract/inter"
)

type Bag map[string]inter.Collection

func NewBag(bags ...map[string][]string) Bag {
	result := Bag{}

	for _, maps := range bags {
		for key, items := range maps {
			for _, item := range items {
				result.Push(key, item)
			}
		}
	}

	return result
}

func NewBagByMap(bags ...Map) Bag {
	result := Bag{}

	for _, items := range bags {
		for key, item := range items {
			result.Push(key, item)
		}
	}

	return result
}

func (b Bag) Source() map[string]inter.Collection {
	return b
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// nil. To access multiple values, use GetCollection or Collection.
func (b Bag) Get(key string) inter.Value {
	values := b[key].Source()
	if len(values) == 0 {
		return nil
	}

	return values[0]
}

func (b Bag) GetMany(key string) inter.Collection {
	if values, found := b[key]; found {
		return values
	}

	return NewCollection()
}

// Set sets the key to value. It replaces any existing
// values.
func (b Bag) Set(key string, value inter.Value) {
	b[key] = NewCollection(value)
}

// Push adds the value to key. It appends to any existing
// values associated with key.
func (b Bag) Push(key string, value interface{}) {
	if collection, found := b[key]; found {
		b[key] = collection.Push(value)
	} else {
		switch value.(type) {
		case Collection:
			b[key] = value.(Collection)
		default:
			b[key] = NewCollection(value)
		}
	}
}

// Del deletes the values associated with key.
func (b Bag) Del(key string) {
	delete(b, key)
}

func (b Bag) Merge(bags ...inter.Bag) inter.Bag {
	for _, bag := range bags {
		for key, item := range bag.Source() {
			b.Push(key, item)
		}
	}

	return b
}
