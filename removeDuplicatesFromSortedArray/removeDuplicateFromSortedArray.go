package removeDuplicatesFromSortedArray

func removeDuplicates(nums []int) int {
	rang := len(nums)-1
	for i:= rang; i > 0; i-- {
		if nums[i] != nums[i-1] {
			nums = append(nums[:])
		} else {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
