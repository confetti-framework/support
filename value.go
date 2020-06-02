package support

import (
	"fmt"
	"strconv"
	"strings"
)

type Value struct {
	source string
	error  error
}

func NewValue(input interface{}) Value {
	var source string

	switch input.(type) {
	case nil:
		source = ""
	case int:
		source = strconv.Itoa(input.(int))
	case string:
		source = input.(string)
	default:
		fmt.Println("type unknown")
	}

	return Value{source: source}
}

func NewValueE(source string, error error) Value {
	return Value{source: source, error: error}
}

func (v Value) String() string {
	if v.error != nil {
		panic(v.error)
	}

	return v.source
}

func (v Value) StringE() (string, error) {
	return v.source, v.error
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

	if v.source == "" {
		return 0, nil
	}

	return strconv.Atoi(v.source)
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
	return strings.Split(v.source, separator)
}
