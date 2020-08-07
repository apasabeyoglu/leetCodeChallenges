package longestCommonPrefix

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	result := strs[0]
	dummy := ""
	for _,str := range strs {
		for j := len(str); j >= 1; j-- {
			dummy = ""
			if strings.Index(result, str[:j]) == 0{
				dummy = str[:j]
				break
			}
		}
		if str != "" {
			result = dummy
		} else {
			return ""
		}
	}
	return result
}