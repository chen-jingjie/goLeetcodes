package arrangeCoins_441

import "math"
/*

*/
func ArrangeCoins(n int) int {
	r := int(math.Floor(math.Sqrt(float64(n * 2))))
	if r * (r + 1) / 2 > n {
		r -= 1
	}

	return r
}
