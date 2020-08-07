package romanToInteger

func romanToInt2(s string) int {
	result := 0

	for i, char := range s {
		dummy := string(char)
		switch dummy {

		}
		if i != len(s)-1 && dummy == "I" && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
			result -= 1
		} else if i != len(s)-1 && dummy == "X" && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
			result -= 10
		} else if i != len(s)-1 && dummy == "C" && (string(s[i+1]) == "D" || string(s[i+1]) == "M"){
			result -= 100
		} else if dummy == "I" {
			result += 1
		} else if dummy == "V" {
			result += 5
		} else if dummy == "X" {
			result += 10
		} else if dummy == "L" {
			result += 50
		} else if dummy == "C" {
			result += 100
		} else if dummy == "D" {
			result += 500
		} else if dummy == "M" {
			result += 1000
		}
	}
	return result
}