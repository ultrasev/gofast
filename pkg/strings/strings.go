package strings

import (
	"strings"
	"unicode"
)

// Reverse returns the reversed string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ToCamelCase converts a string to camelCase
// Example: "hello_world" -> "helloWorld"
func ToCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for i := 1; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

// ToKebabCase converts a string to kebab-case
// Example: "helloWorld" -> "hello-world"
func ToKebabCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result.WriteRune('-')
		}
		result.WriteRune(unicode.ToLower(r))
	}
	return result.String()
}

// Truncate returns a truncated string with optional suffix
func Truncate(s string, length int, suffix string) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + suffix
}

// IsEmpty checks if string is empty or contains only whitespace
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// RemoveSpaces removes all whitespace from string
func RemoveSpaces(s string) string {
	return strings.Join(strings.Fields(s), "")
}

// ContainsAny checks if string contains any of the given substrings
func ContainsAny(s string, substrings ...string) bool {
	for _, sub := range substrings {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}
