package util

import "strings"

func UnderscoreToCamelCase(s string) string {
	parts := StringParts(strings.Split(s, "_"))
	parts[1:].Each(strings.Title)
	return strings.Join(parts, "")
}

func LowercaseFirst(s string) string {
	return strings.ToLower(s[0:1]) + s[1:]
}
