package support

import (
	"reflect"
	"strconv"
)

type Collection []Value

func NewCollection(items ...interface{}) Collection {
	collection := Collection{}

	for _, item := range items {
		if inputCollection, ok := item.(Collection); ok {
			collection = append(collection, inputCollection...)
			continue
		}

		switch Type(item) {
		case reflect.Array, reflect.Slice:
			s := reflect.ValueOf(item)
			for i := 0; i < s.Len(); i++ {
				value := s.Index(i).Interface()
				collection = append(collection, NewValue(value))
			}
		default:
			collection = append(collection, NewValue(item))
		}
	}

	return collection
}

func (c Collection) Raw() interface{} {
	var result []interface{}
	var raw interface{}

	for _, value := range c {
		raw = value.Raw()
		result = append(result, raw)
	}

	return result
}

func (c Collection) Get(key string) Value {
	result, err := c.GetE(key)
	if err != nil {
		panic(err)
	}
	return result
}

func (c Collection) GetE(key string) (Value, error) {
	if key == "" {
		return NewValue(c), nil
	}

	currentKey, rest := splitKey(key)

	// when you request something with an Asterisk, you always develop a collection
	if currentKey == "*" {

		flattenCollection := Collection{}
		for _, value := range c {
			switch Type(value.Source()) {
			case reflect.Slice, reflect.Array:
				flattenCollection = append(flattenCollection, value.Source().(Collection)...)
			case reflect.Map:
				flattenCollection = append(flattenCollection, value.Source().(Map).Collection()...)
			default:
				return NewValue(c), nil
			}
		}

		return flattenCollection.GetE(joinRest(rest))
	}

	index, err := strconv.Atoi(key)
	if err != nil {
		return Value{}, InvalidCollectionKeyError.Wrap("'%s' can only be a number or *", key)
	}

	if len(c) < (index + 1) {
		return Value{}, CanNotFoundValueError.Wrap("'%s' not found", key)
	}

	return c[index], nil
}

func (c Collection) First() Value {
	if len(c) == 0 {
		return Value{}
	}

	return c[0]
}

func (c Collection) Push(item interface{}) Collection {
	return append(c, NewValue(item))
}

func (c Collection) SetE(key string, value interface{}) (Collection, error) {
	if key == "" {
		return c.Push(value), nil
	}

	currentKey, rest := splitKey(key)

	if len(rest) == 0 {
		return c.Push(value), nil
	}
	var nestedValue interface{}
	var err error

	if currentKey == "*" {
		nestedValue, err := NewValue(nil).SetE(joinRest(rest), value)
		if err != nil {
			return c, err
		}
		c.Push(nestedValue)
	}

	nestedValue, err = NewMap().SetE(joinRest(rest), value)
	if err != nil {
		return nil, err
	}
	return c.Push(nestedValue), nil
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
