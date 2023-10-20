// Package snakecase provides a utility to convert strings into snake_case.
package snakecase

import (
	"regexp"
	"strings"
)

// ToSnakeCase converts a string into snake_case.
func ToSnakeCase(s string) string {
	// This regex will match any uppercase letter and ensures they are separated correctly.
	// We will use two regex patterns: one to identify the uppercase letters preceded by lowercase or numbers,
	// and one to identify consecutive uppercase letters.
	s = regexp.MustCompile("([a-z0-9]+|[A-Z])([A-Z])").ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(s)
}
