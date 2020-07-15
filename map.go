package support

import (
	"errors"
	"net/url"
)

type Map map[string]Value

func NewMap(itemsRange ...map[string]interface{}) Map {
	result := Map{}

	for _, items := range itemsRange {
		for key, item := range items {
			result[key] = NewValue(item)
		}
	}

	return result
}

func NewMapByString(itemsRange ...map[string]string) Map {
	result := Map{}

	for _, items := range itemsRange {
		for key, item := range items {
			result[key] = NewValue(item)
		}
	}

	return result
}

func NewMapByUrlValues(itemsRange ...url.Values) Map {
	result := Map{}

	for _, items := range itemsRange {
		for key, strings := range items {
			collection := NewCollection()
			for _, stringItem := range strings {
				collection = collection.Push(stringItem)
			}
			result[key] = NewValue(collection)
		}
	}

	return result
}

func (m Map) Source() map[string]Value {
	return m
}

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// nil. To access multiple values, use GetCollection or Collection.
func (m Map) Get(key string) Value {
	if key == "" {
		return NewValue(m)
	}

	currentKey, rest := splitKey(key)

	// when you request something with an Asterisk, you always develop a collection
	if currentKey == "*" {
		collection := NewCollection()
		for _, values := range m.Source() {
			for _, value := range values.Collection() {
				collection = collection.Push(value)
			}
		}

		return NewValue(collection)
	}

	value, found := m[currentKey]

	if !found {
		return NewValueE(nil, errors.New(key + " not found"))
	}

	switch value.Source().(type) {
	case Collection:
		return value.Collection().Get(joinRest(rest))
	case Map:
		deeperValue, found := value.Source().(Map)[key]
		if found {
			return deeperValue
		}

		return NewValueE(nil, errors.New("no value found with key "+key))
	default:
		return value
	}
}

// Set sets the key to value. It replaces any existing
// values.
func (m Map) Set(key string, value Value) Map {
	m[key] = value

	return m
}

// Push adds the value to key. It appends to any existing values
// associated with key. If the value is in collection, push
// the value to the collection.
func (m Map) Push(key string, input interface{}) Map {
	if rawValue, found := m[key]; found {
		source := rawValue.Source()
		switch source.(type) {
		case Collection:
			collection := source.(Collection)
			m[key] = NewValue(collection.Push(input))
		default:
			m[key] = NewValue(input)
		}
	} else {
		m[key] = NewValue(input)
	}

	return m
}

// Delete deletes the values associated with key.
func (m Map) Delete(key string) {
	delete(m, key)
}

func (m Map) Collection() Collection {
	collection := NewCollection()
	for _, value := range m {
		collection = collection.Push(value)
	}

	return collection
}

func (m Map) Merge(maps ...Map) Map {
	for _, bag := range maps {
		for key, item := range bag.Source() {
			m.Push(key, item)
		}
	}

	return m
}

func (m Map) First() Value {
	return m.Collection().First()
}

func (m Map) Empty() bool {
	return len(m) == 0
}
