package support

import (
	"errors"
	"reflect"
)

type Map map[string]Value

func NewMap(itemsRange ...interface{}) Map {
	result := Map{}

	for _, rawItems := range itemsRange {
		v := reflect.ValueOf(rawItems)
		if v.Kind() != reflect.Map {
			panic("can't create map from " + v.Kind().String())
		}

		for _, key := range v.MapKeys() {
			result[key.String()] = NewValue(v.MapIndex(key).Interface())
		}
	}

	return result
}

func (m Map) Raw() interface{} {
	result, err := m.RawE()
	if err != nil {
		panic(err)
	}

	return result
}

func (m Map) RawE() (interface{}, Errors) {
	result := map[string]interface{}{}
	var err Errors

	for key, value := range m {
		raw, valErr := value.RawE()

		// Handle value
		result[key] = raw

		// Handle errors
		if multiErr, ok := valErr.(Errors); ok {
			err = append(err, multiErr...)
		} else if valErr != nil {
			err = append(err, valErr)
		}
	}

	return result, err
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
		for _, values := range m {
			for _, value := range values.Collection() {
				collection = collection.Push(value)
			}
		}

		return NewValue(collection)
	}

	value, found := m[currentKey]
	if !found {
		return NewValueE(nil, errors.New(key+" not found"))
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

func (m Map) Only(keys ...string) Map {
	result := Map{}
	for _, key := range keys {
		if m.Has(key) {
			result.Set(key, m.Get(key))
		}
	}

	return result
}

func (m Map) Except(keys ...string) Map {
	result := m.Copy()
	for _, key := range keys {
		delete(result, key)
	}

	return result
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
		for key, item := range bag {
			m.Push(key, item)
		}
	}

	return m
}

// Generates a new struct with the same data as the old struct
func (m Map) Copy() Map {
	newMap := Map{}
	for key, value := range m {
		newMap[key] = value
	}

	return newMap
}

func (m Map) First() Value {
	return m.Collection().First()
}

func (m Map) Has(keys ...string) bool {
	for _, key := range keys {
		// todo don't check for nil error, but for NotFound error
		if m.Get(key).Error() != nil {
			return false
		}
	}

	return true
}

func (m Map) HasAny(keys ...string) bool {
	if len(keys) == 0 {
		return true
	}

	for _, key := range keys {
		// todo don't check for nil error, but for NotFound error
		if m.Get(key).Error() == nil {
			return true
		}
	}

	return false
}

func (m Map) Missing(keys ...string) bool {
	return !m.Has(keys...)
}

func (m Map) Filled(keys ...string) bool {
	for _, key := range keys {
		if m.Get(key).Empty() {
			return false
		}
	}

	return true
}

func (m Map) Empty() bool {
	return len(m) == 0
}
