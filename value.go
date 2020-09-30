package support

import (
	"errors"
	"github.com/spf13/cast"
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
		return Value{source: NewMap(value)}
	}

	return Value{source: value}
}

func NewValueE(val interface{}, unknownError interface{}) Value {
	var err error

	switch knownError := unknownError.(type) {
	case string:
		err = errors.New(knownError)
	case error:
		err = knownError
	case nil:
		err = nil
	default:
		panic("can't convert variable in an error (type " + Name(knownError) + ")")
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
	nextKey := joinRest(rest)

	// when you request something with an Asterisk, you always develop a collection
	if currentKey == "*" {

		switch v.source.(type) {
		case Collection:
			return v.source.(Collection).Get(nextKey)
		case Map:
			return v.source.(Map).Get(nextKey)
		default:
			return NewValueE(nil, errors.New("*: is not a Collection or Map"))
		}

	}

	switch source := v.source.(type) {
	case Collection:
		keyInt, err := strconv.Atoi(currentKey)
		if err != nil {
			return NewValueE(nil, err)
		}
		return v.source.(Collection)[keyInt].Get(nextKey)
	case Map:
		return v.source.(Map)[currentKey].Get(nextKey)
	default:
		switch Type(source) {
		case reflect.Struct:
			val := reflect.ValueOf(source).FieldByName(currentKey)
			if val.IsValid() {
				return NewValue(val.Interface()).Get(nextKey)
			} else {
				return NewValueE(nil, errors.New(currentKey+": can't find value"))
			}

		}
		return NewValueE(nil, errors.New(currentKey+": is not a struct, Collection or Map"))
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
	switch source := v.source.(type) {
	case Collection:
		result, err = source.First().StringE()
	case Map:
		result, err = source.First().StringE()
	default:
		result, err = cast.ToStringE(source)
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
	switch source := v.source.(type) {
	case Collection:
		result, err = source.First().NumberE()
	case Map:
		result, err = source.First().NumberE()
	default:
		result, err = cast.ToIntE(source)
	}

	if v.error != nil {
		err = v.error
	}

	return
}

func (v Value) Float() float64 {
	values, err := v.FloatE()
	if err != nil {
		panic(err)
	}

	return values
}

func (v Value) FloatE() (result float64, err error) {

	switch source := v.source.(type) {
	case Collection:
		result, err = source.First().FloatE()
	case Map:
		result, err = source.First().FloatE()
	default:
		result, err = cast.ToFloat64E(source)
	}

	if v.error != nil {
		err = v.error
	}

	return
}

func (v Value) Bool() bool {
	result, err := v.BoolE()
	if err != nil {
		panic(err)
	}

	return result
}

func (v Value) BoolE() (bool, error) {
	result, err := v.RawE()

	switch result {
	case true, 1, "1", "true", "on", "yes":
		return true, err
	default:
		return false, err
	}
}

func (v Value) Filled() bool {
	return v.source != nil && v.source != ""
}

func (v Value) Empty() bool {
	return v.source == nil || v.source == ""
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
