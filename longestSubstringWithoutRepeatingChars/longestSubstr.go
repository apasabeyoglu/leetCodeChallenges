package longestSubstringWithoutRepeatingChars

import "strings"

func lengthOfLongestSubstring(s string) int {
	result := ""
	resultFinal := ""

	for i,char := range s {
		index := strings.Index(result,string(char))
		if index == -1 {
			result += string(char)
			if i == len(s)-1 && len(result) > len(resultFinal) {
				return len(result)
			}
		} else {
			if len(result) > len(resultFinal) {
				resultFinal = result
			}
			result = result[index+1:] + string(char)
		}
	}
	return len(resultFinal)
}
