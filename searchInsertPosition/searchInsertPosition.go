package searchInsertPosition

import "sort"

func searchInsert(nums []int, target int) int {

	for i, num := range nums {
		if target == num || target < num {
			return i
		}
		if i == len(nums)-1 && target > num {
			return len(nums)
		}
		if target < nums[i+1] && target > num {
			return i+1
		}
	}
	return 0
}

func searchInsert2(nums []int, target int) int {

	nums = append(nums, target)
	sort.Ints(nums)

	for i,num := range nums {
		if target == num {
			return i
		}
	}

	return 0
}
