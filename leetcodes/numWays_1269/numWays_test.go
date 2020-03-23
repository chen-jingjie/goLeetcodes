package numWays_1269

import (
	"fmt"
	"testing"
)

func TestNumWays(t *testing.T) {
	values := [][]int{
		{3, 2},
	}
	results := []int{4, 2, 8}

	for i, value := range values {
		if v := NumWays(value[0],value[1]); v != results[i] {
			t.Error(fmt.Sprintf("错误：steps %v, arrLen: %v, need: %v, got: %v",
				value[0], value[1], results[i], v))
		}
	}
}
