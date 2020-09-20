package removeElement

func removeElement(nums []int, val int) int {
	rang := len(nums)-1

	for i := rang; i >= 0; i-- {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}