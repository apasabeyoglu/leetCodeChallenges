package lengthOfLastWord

import "strings"

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	words := strings.Split(s, " ")

	for i := len(words)-1; i >= 0; i-- {
		if words[i] != "" {
			return len(words[i])
		}
	}
	return 0
}
