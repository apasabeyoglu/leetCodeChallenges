package validParentheses

import "strings"

func isValid(s string) bool {
	for i := len(s); i >= 2; i--{
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
	}
	if len(s) != 0 {
		return false
	}
	return true
}
