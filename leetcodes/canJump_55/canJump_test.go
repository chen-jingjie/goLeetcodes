package canJump_55

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	values := [][]int{
		{2, 3, 1, 1, 4},
		{3, 2, 1, 0, 4},
		{0},
	}
	results := []bool{true, false, true}

	for i, value := range values {
		if CanJump(value) != results[i] {
			fmt.Println(value, CanJump(value))
			t.Error("错误")
		}
	}
}
