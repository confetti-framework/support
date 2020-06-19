package support

import (
	"errors"
	"github.com/lanvard/contract/inter"
)

type Collection []inter.Value

func NewCollection(items ...interface{}) Collection {
	collection := Collection{}

	for _, item := range items {
		collection = append(collection, NewValue(item))
	}

	return collection
}

func (c Collection) Source() []inter.Value {
	return c
}

func (c Collection) First() inter.Value {
	if len(c) == 0 {
		return NewValueE("", errors.New("value not found in collection"))
	}

	return c[0]
}

func (c Collection) Push(item interface{}) inter.Collection {
	result := append(c, NewValue(item))

	return result
}

func (c Collection) Reverse() inter.Collection {
	items := c
	for left, right := 0, len(items)-1; left < right; left, right = left+1, right-1 {
		items[left], items[right] = items[right], items[left]
	}

	return items
}

// Determine if an item exists in the collection by a string
func (c Collection) Contains(search interface{}) bool {
	for _, item := range c {
		if item.Source() == search {
			return true
		}
	}

	return false
}

// The len method returns the length of the collection
func (c Collection) Len() int {
	return len(c.Source())
}
