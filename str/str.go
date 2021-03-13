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

// Return the remainder of a string after the first occurrence of a given value.
func After(subject string, search string) string {
	l := len(search)
	if l == 0 {
		return ""
	}

	byteIndex := strings.Index(subject, search)
	if byteIndex == -1 {
		return subject
	}

	byteSubject := []byte(subject)
	byteSearch := []byte(search)

	result := string(byteSubject[byteIndex+ len(byteSearch):])
	return result
}

// Return the remainder of a string after the last occurrence of a given value.
func AfterLast(subject string, search string) string {
	l := len(search)
	if l == 0 {
		return ""
	}

	byteIndex := strings.LastIndex(subject, search)
	if byteIndex == -1 {
		return subject
	}

	byteSubject := []byte(subject)
	byteSearch := []byte(search)

	result := string(byteSubject[byteIndex+ len(byteSearch):])
	return result
}

// Get the portion of a string before the first occurrence of a given value.
func Before(subject string, search string) string {
	l := len(search)
	if l == 0 {
		return subject
	}

	byteIndex := strings.Index(subject, search)
	if byteIndex == -1 {
		return subject
	}

	byteSubject := []byte(subject)

	result := string(byteSubject[:byteIndex])
	return result
}

func BeforeLast(subject string, search string) string {
	l := len(search)
	if l == 0 {
		return subject
	}

	byteIndex := strings.LastIndex(subject, search)
	if byteIndex == -1 {
		return subject
	}

	byteSubject := []byte(subject)

	result := string(byteSubject[:byteIndex])
	return result
}

func Between(subject string, from string, to string) string {
	if len(from) == 0 || len(to) == 0 {
		return subject
	}
	return BeforeLast(After(subject, from), to)
}

func Contains(haystack string, needle string) bool {
	return strings.Index(haystack, needle) == -1
}

func ContainsFromSlice(haystack string, needle []string) bool {
	// TODO
	return false
}

func ContainsAllFromSlice(haystack string, needle []string) bool {
	// TODO
	return false
}

func EndsWith(haystack string, needle string) bool {
	// TODO
	return false
}

func Finish(value string, cap string) string {
	// TODO
	return ""
}

func Kebab(vale string) string {
	// TODO
	return ""
}

func Length(value string) int {
	// TODO
	return 0
}

func LimitCharacters(value string, limit int, end string) string{
	// TODO
	return ""
}

func LimitWords(value string, limit int, end string) string{
	// TODO
	return ""
}

func Lower(value string) string {
	// TODO
	return ""
}

func PadBoth(value string, length int, pad string) string {
	// TODO
	return ""
}

func PadLeft(value string, length int, pad string) string {
	// TODO
	return ""
}

func PadRight(value string, length int, pad string) string {
	// TODO
	return ""
}

func ReplaceArray(search string, replace []string, subject string) string {
	// TODO
	return ""
}

func ReplaceFirst(search string, replace string, subject string) string {
	// TODO
	return ""
}

func ReplaceLast(search string, replace string, subject string) string {
	// TODO
	return ""
}

func Start(value string, prefix string) string {
	// TODO
	return ""
}

func Slug(value string) string {
	// TODO
	return ""
}
func SlugWithDelimiter(value string, delimiter string) string {
	// TODO
	return ""
}

func Snake(value string) string {
	// TODO
	return ""
}

func SnakeWithDelimiter(value string, delimiter string) string {
	// TODO
	return ""
}

func StartsWith(haystack string, needle string) string {
	// TODO
	return ""
}

func Studly(value string) string {
	// TODO
	return ""
}

func UcFirst(value string) string {
	// TODO
	return ""
}

func Upper(value string) string {
	// TODO
	return ""
}

func Title(value string) string {
	// TODO
	return ""
}





















