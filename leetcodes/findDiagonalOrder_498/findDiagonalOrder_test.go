package findDiagonalOrder_498

import "testing"

func TestFindDiagonalOrder(t *testing.T) {
	v := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	r := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}
	computed := FindDiagonalOrder(v)

	for i, val := range r {
		if computed[i] != val {
			t.Error("错误")
		}
	}
}
