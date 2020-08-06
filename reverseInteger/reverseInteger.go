package reverseInteger

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	result := ""
	isNegative := false

	if x<0 {
		isNegative = true
		x *= -1
	}

	dummyStr := strconv.Itoa(x)
	for _,char := range dummyStr {
		result = string(char) + result
	}

	resultInt,_ := strconv.Atoi(result)

	if resultInt > math.MaxInt32 {
		return 0
	}
	if isNegative {
		resultInt *= -1
	}
	return resultInt
}
