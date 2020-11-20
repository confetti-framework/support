package support

import (
	"reflect"
)

type Map map[string]Value

func NewMap(itemsRange ...interface{}) Map {
	result, err := NewMapE(itemsRange...)
	if err != nil {
		panic(err)
	}
	return result
}

func NewMapE(itemsRange ...interface{}) (Map, error) {
	var err error
	result := Map{}

	for _, rawItems := range itemsRange {
		v := reflect.ValueOf(rawItems)
		if v.Kind() != reflect.Map {
			err = CanNotCreateMapError.Wrap("type %s", v.Kind().String())
			continue
		}

		for _, key := range v.MapKeys() {
			value := v.MapIndex(key).Interface()
			result[key.String()] = NewValue(value)
		}
	}

	return result, err
}

func (m Map) Raw() interface{} {
	result := map[string]interface{}{}

	for key, value := range m {
		// Handle value
		result[key] = value.Raw()
	}

	return result
}

func (m Map) Get(key string) Value {
	result, err := m.GetE(key)
	if err != nil {
		panic(err)
	}
	return result
}

// GetE gets the first value associated with the given key.
// If there are no values associated with the key, GetE returns
// nil. To access multiple values, use GetCollection or Collection.
func (m Map) GetE(key string) (Value, error) {
	if key == "" {
		return NewValue(m), nil
	}

	currentKey, rest := splitKey(key)

	// when you request something with an Asterisk, you always develop a collection
	if currentKey == "*" {
		collection := Collection{}
		for _, values := range m {
			for _, value := range values.Collection() {
				collection = collection.Push(value)
			}
		}

		return NewValue(collection), nil
	}

	value, found := m[currentKey]
	if !found {
		return Value{}, CanNotFoundValueError.Wrap("key '%s'%s", currentKey, getKeyInfo(key, currentKey))
	}

	switch value.Source().(type) {
	case Collection:
		return value.Collection().GetE(joinRest(rest))
	case Map:
		return value.Map().GetE(joinRest(rest))
	default:
		return value.GetE(joinRest(rest))
	}
}

func getKeyInfo(key string, currentKey string) string {
	info := ""
	if currentKey != key {
		info = " ('" + key + "')"
	}
	return info
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
		item, err := m.GetE(key)
		if err == nil {
			result.Set(key, item)
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
	collection := Collection{}
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
		_, err := m.GetE(key)
		if err != nil {
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
		result, err := m.GetE(key)
		if err != nil && !result.Empty() {
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
		result, err := m.GetE(key)
		if err != nil || result.Empty() {
			return false
		}
	}

	return true
}

func (m Map) Empty() bool {
	return len(m) == 0
}
