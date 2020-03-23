package findDiagonalOrder_498

import "fmt"

func FindDiagonalOrder(matrix [][]int) []int {
	direction := true
	var m, n, mMax, nMax, pre int

	mMax = len(matrix)
	if mMax == 0 {
		return []int{}
	}

	nMax = len(matrix[0])
	if nMax == 0 {
		return []int{}
	}

	r := make([]int, mMax*nMax)

	mMax -= 1
	nMax -= 1

	for m <= mMax || n <= nMax {
		fmt.Println(m, n, pre)

		r[pre] = matrix[m][n]
		pre++
		if m == mMax && n == nMax {
			return r
		}
		if direction {
			m -= 1
			n += 1
		} else {
			m += 1
			n -= 1
		}

		if m < 0 {
			m = 0
			direction = !direction

			if n > nMax {
				n = nMax
				m = 1
			}
		} else if m <= mMax {
			if n < 0 {
				n = 0
				direction = !direction
			} else if n > nMax {
				n = nMax
				m += 2
				direction = !direction
			}
		} else {
			m = mMax
			n += 2
			direction = !direction
		}
	}

	return r
}
