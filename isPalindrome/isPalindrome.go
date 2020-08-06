package isPalindrome

func isPalindrome(x int) bool {
	if x<0 {
		return false
	}

	divisor := 1

	for x/divisor >= 10{
		divisor *= 10
	}

	for x > 0 {
		lead := x/divisor
		trail := x%10
		if lead != trail {
			return false
		}
		x = (x % divisor)/10

		divisor = divisor / 100
	}
	return true
}
