package support

import (
	"errors"
	"strconv"
	"strings"
)

type Value struct {
	source interface{}
	error  error
}

func NewValue(val interface{}) Value {
	switch val.(type) {
	case Value:
		return val.(Value)
	default:
		return Value{source: val}
	}
}

func NewValueE(val interface{}, error error) Value {
	switch val.(type) {
	case Value:
		return Value{source: val.(Value).source, error: error}
	default:
		return Value{source: val, error: error}
	}
}

func (v Value) Source() interface{} {
	return v.source
}

func (v Value) String() string {
	result, err := v.StringE()
	if err != nil {
		panic(v.error)
	}

	return result
}

func (v Value) StringE() (string, error) {
	if v.error != nil {
		return "", v.error
	}

	switch v.source.(type) {
	case nil:
		return "", nil
	case int:
		return strconv.Itoa(v.source.(int)), nil
	case string:
		return v.source.(string), nil
	default:
		return "", errors.New("type unknown")
	}
}

func (v Value) Strings() []string {
	values, err := v.StringsE()
	if err != nil {
		panic(err)
	}

	return values
}

func (v Value) StringsE() ([]string, error) {
	return v.Split(","), v.error
}

func (v Value) Number() int {
	values, err := v.NumberE()
	if err != nil {
		panic(err)
	}

	return values
}

func (v Value) NumberE() (int, error) {
	if v.error != nil {
		return 0, v.error
	}

	if v.String() == "" {
		return 0, nil
	}

	return strconv.Atoi(v.String())
}

func (v Value) Numbers() []int {
	values, err := v.NumbersE()
	if err != nil {
		panic(err)
	}

	return values
}

func (v Value) NumbersE() ([]int, error) {
	rawValues := v.Split(",")
	if v.error != nil {
		return nil, v.error
	}

	var result []int

	for _, rawValue := range rawValues {
		value, err := strconv.Atoi(rawValue)
		if err != nil {
			return nil, err
		}

		result = append(result, value)
	}

	return result, nil
}

func (v Value) Empty() bool {
	return "" == v.source
}

func (v Value) Present() bool {
	return "" != v.source
}

// Split slices Value into all substrings separated by separator and returns a slice of
// the strings between those separators.
//
// If Value does not contain separator and separator is not empty, Split returns a
// slice of length 1 whose only element is Value.
func (v Value) Split(separator string) []string {
	return strings.Split(v.String(), separator)
}
