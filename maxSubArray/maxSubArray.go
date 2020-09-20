package maxSubArray

//Kadane's algorithm
func maxSubArray(nums []int) int {
	curr := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > curr + nums[i] {
			curr = nums[i]
		} else {
			curr = curr + nums[i]
		}
		if result < curr {
			result = curr
		}
	}
	return result
}
