package twoSum_1

func TwoSum(nums []int, target int) []int {
	v := make(map[int]int)
	for i := range nums {
		val := target - nums[i]
		c, ok := v[nums[i]]
		if ok != false {
			return []int{c, i}
		}
		v[val] = i
	}
	return []int{-1, -1}
}
