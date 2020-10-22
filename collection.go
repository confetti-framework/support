package support

import (
	"errors"
	"reflect"
	"strconv"
)

type Collection []Value

func NewCollection(items ...interface{}) Collection {
	collection, err := NewCollectionE(items...)
	if err != nil {
		panic(err)
	}
	return collection
}

func NewCollectionE(items ...interface{}) (Collection, error) {
	collection := Collection{}
	var err error

	for _, item := range items {
		if inputCollection, ok := item.(Collection); ok {
			collection = append(collection, inputCollection...)
			continue
		}

		switch Type(item) {
		case reflect.Array, reflect.Slice:
			if interfaces, ok := item.([]interface{}); ok {
				for _, value := range interfaces {
					collection = append(collection, NewValue(value))
				}
			} else if strings, ok := item.([]string); ok {
				for _, value := range strings {
					collection = append(collection, NewValue(value))
				}
			} else {
				v := reflect.ValueOf(item)
				err = errors.New("Can't create collection from type " + v.Kind().String())
			}
		default:
			collection = append(collection, NewValue(item))
		}
	}

	return collection, err
}

func (c Collection) Raw() interface{} {
	collection, err := c.RawE()
	if err != nil {
		panic(err)
	}

	return collection
}

func (c Collection) RawE() (interface{}, error) {
	var result []interface{}
	var raw interface{}
	var err error

	for _, value := range c {
		raw, err = value.RawE()
		result = append(result, raw)
	}

	return result, err
}

func (c Collection) Get(key string) Value {
	if key == "" {
		return NewValue(c)
	}

	currentKey, rest := splitKey(key)

	// when you request something with an Asterisk, you always develop a collection
	if currentKey == "*" {

		flattenCollection := NewCollection()
		for _, value := range c {
			switch Type(value.Source()) {
			case reflect.Slice, reflect.Array:
				flattenCollection = append(flattenCollection, value.Source().(Collection)...)
			case reflect.Map:
				flattenCollection = append(flattenCollection, value.Source().(Map).Collection()...)
			case reflect.String, reflect.Int, reflect.Float64, reflect.Float32:
				return NewValue(c)
			default:
				panic("value " + Name(value) + " has a unknown type ")
			}
		}

		return NewValue(flattenCollection.Get(joinRest(rest)))
	}

	index, err := strconv.Atoi(key)
	if err != nil {
		return NewValueE(nil, errors.New(key+" can only be a number or *"))
	}

	if len(c) < (index + 1) {
		return NewValueE(nil, errors.New(key+" not found"))
	}

	return c[index]
}

func (c Collection) First() Value {
	if len(c) == 0 {
		return NewValueE("", errors.New("value not found in collection"))
	}

	return c[0]
}

func (c Collection) Push(item interface{}) Collection {
	return append(c, NewValue(item))
}

func (c Collection) Reverse() Collection {
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
	return len(c)
}

func (c Collection) Empty() bool {
	return len(c) == 0
}
