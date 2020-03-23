package canJump_55

func CanJump(nums []int) bool {
	l := len(nums)
	good := l - 1
	for i := l - 2; i >= 0; i-- {
		if nums[i] >= good-i {
			good = i
		}
	}
	return good == 0
}

// 上面优化一版
//func CanJump(nums []int) bool {
//	l := len(nums)
//	if l == 1 {
//		return true
//	}
//	for i := 0; i < l-1; i++ {
//		if nums[i] >= l - 1 - i {
//			if i == 0 {
//				return true
//			}
//			return CanJump(nums[:i+1])
//		}
//	}
//	return false
//}
