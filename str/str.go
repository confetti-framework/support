package str

import (
	"strings"
	"unicode"
)

func UpperFirst(input string) string {
	if len(input) == 0 {
		return ""
	}
	tmp := []rune(input)
	tmp[0] = unicode.ToUpper(tmp[0])
	return string(tmp)
}

func InSlice(input interface{}, expects ...interface{}) bool {
	for _, expect := range expects {
		if input == expect {
			return true
		}
	}
	return false
}

func After(input string, search string) string {
	if len(search) == 0 {
		return input
	}
	results := strings.SplitN(input, search, 2)
	return results[len(results)-1]
}

func AfterLast(input string, search string) string {
	if len(search) == 0 {
		return input
	}
	position := strings.LastIndex(input, search)

	if position == -1 {
		return input
	}

	return input[position+len(search):]
}

// TBD: Ascii

func Before(input string, search string) string {
	if len(search) == 0 {
		return input
	}
	position := strings.Index(input, search)

	if position == -1 {
		return input
	}

	return input[:position]
}

func BeforeLast(input string, search string) string {
	if len(search) == 0 {
		return input
	}
	position := strings.LastIndex(input, search)

	if position == -1 {
		return input
	}

	return input[:position]
}