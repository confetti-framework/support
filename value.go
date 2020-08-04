package support

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Value struct {
	source interface{}
	error  error
}

func NewValue(value interface{}) Value {
	switch value.(type) {
	case Collection, Map:
		return Value{source: value}
	case Value:
		return value.(Value)
	}

	switch Type(value) {
	case reflect.Slice, reflect.Array:
		return NewValue(NewCollection(value))
	case reflect.Map:
		return Value{source: NewMap(value.(map[string]interface{}))}
	}

	return Value{source: value}
}

func NewValueE(val interface{}, inputErr interface{}) Value {
	var err error
	if e, ok := inputErr.(string); ok {
		err = errors.New(e)
	}

	switch val.(type) {
	case Value:
		return Value{source: val.(Value).source, error: err}
	default:
		return Value{source: val, error: err}
	}
}

func (v Value) Source() interface{} {
	return v.source
}

func (v Value) Raw() interface{} {
	value, err := v.RawE()
	if err != nil && err.Error() != "" {
		panic(err)
	}

	return value
}

func (v Value) RawE() (interface{}, error) {
	if result, ok := v.source.(Value); ok {
		return result.RawE()
	}
	if result, ok := v.source.(Collection); ok {
		return result.RawE()
	}
	if result, ok := v.source.(Map); ok {
		return result.RawE()
	}

	return v.source, v.error
}

func (v Value) Error() error {
	return v.error
}

func (v Value) Get(key string) Value {
	if key == "" || v.Error() != nil {
		return v
	}

	currentKey, rest := splitKey(key)

	// when you request something with an Asterisk, you always develop a collection
	if currentKey == "*" {

		switch v.source.(type) {
		case Collection:
			return v.source.(Collection).Get(joinRest(rest))
		case Map:
			return v.source.(Map).Get(joinRest(rest))
		default:
			return NewValueE(nil, errors.New("*: is not a Collection or Map"))
		}

	}

	switch v.source.(type) {
	case Collection:
		keyInt, err := strconv.Atoi(currentKey)
		if err != nil {
			return NewValueE(nil, err)
		}
		return v.source.(Collection)[keyInt].Get(joinRest(rest))
	case Map:
		return v.source.(Map)[currentKey].Get(joinRest(rest))
	default:
		return NewValueE(nil, errors.New(currentKey+": is not a Collection or Map"))
	}
}

// A value can contain a collection.
func (v Value) Collection() Collection {
	switch v.source.(type) {
	case Collection:
		return v.source.(Collection)
	case Map:
		return v.source.(Map).Collection()
	default:
		return NewCollection(v.source)
	}
}

// A value can contain a Map.
func (v Value) Map() Map {
	switch valueType := v.source.(type) {
	case Map:
		return v.source.(Map)
	default:
		panic("can't create map from reflect.Kind " + strconv.Itoa(int(Type(valueType))))
	}
}

func (v Value) String() string {
	result, err := v.StringE()
	if err != nil {
		panic(err)
	}

	return result
}

func (v Value) StringE() (result string, err error) {
	switch v.source.(type) {
	case nil:
		result = ""
	case int:
		result = strconv.Itoa(v.source.(int))
	case float32:
		result = strconv.FormatFloat(float64(v.source.(float32)), 'E', -1, 32)
	case float64:
		result = strconv.FormatFloat(v.source.(float64), 'f', -1, 64)
	case string:
		result = v.source.(string)
	case Collection:
		result, err = v.Collection().First().StringE()
	case Map:
		return v.Map().First().StringE()
	default:
		err = errors.New("can't convert value to string")
	}

	if v.error != nil {
		err = v.error
	}

	return
}

func (v Value) Number() int {
	values, err := v.NumberE()
	if err != nil {
		panic(err)
	}

	return values
}

func (v Value) NumberE() (result int, err error) {
	switch v.source.(type) {
	case int:
		result = v.source.(int)
	case string:
		stringValue := v.source.(string)
		if stringValue == "" {
			result = 0
		} else {
			result, err = strconv.Atoi(stringValue)
		}
	case Collection:
		result, err = v.Collection().First().NumberE()
	default:
		err = errors.New("can't convert value to number")
	}

	if v.error != nil {
		err = v.error
	}

	return
}

func (v Value) Empty() bool {
	return v.source == nil || v.source == ""
}

func (v Value) Present() bool {
	return v.source != nil && v.source != ""
}

// Split slices Value into all substrings separated by separator and returns a slice of
// the strings between those separators.
//
// If Value does not contain separator and separator is not empty, Split returns a
// slice of length 1 whose only element is Value.
func (v Value) Split(separator string) Collection {
	rawStrings := strings.Split(v.String(), separator)
	var result Collection
	for _, rawString := range rawStrings {
		result = append(result, NewValue(rawString))
	}

	return NewCollection(result)
}
