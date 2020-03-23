package climbStairs_70

func ClimbStairs(n int) int {
	v := make(map[int]int)
	v[1] = 1
	v[2] = 2

	for i := 3; i <= n; i++ {
		v[i] = v[i-1] + v[i-2]
		delete(v, i-2)
	}

	return v[n]
}

// 先用了递归，但是递归时间复杂度太高了，上面优化一下
//func ClimbStairs(n int) int {
//	switch n {
//		case 0:
//			return 0
//		case 1:
//			return 1
//		case 2:
//			return 2
//		default:
//			return ClimbStairs(n-1) + ClimbStairs(n-2)
//	}
//}
