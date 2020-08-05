package twoSum

func twoSum(nums []int, target int) []int {

	mp := make(map[int]int)

	for i, num := range nums {
		a, ok := mp[target-num]
		if ok {
			return []int{a,i}
		}
		mp[nums[i]] = i
	}
	return nil
}