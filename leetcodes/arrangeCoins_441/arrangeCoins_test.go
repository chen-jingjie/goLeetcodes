package arrangeCoins_441

import (
	"fmt"
	"testing"
)

func TestArrangeCoins(t *testing.T) {
	vs := []int{5, 8, 0, 9, 10, 2146467959}
	r := []int{2, 3, 0, 3, 4, 65519}
	for i, v := range vs {
		if rv := ArrangeCoins(v); rv != r[i] {
			t.Error(fmt.Sprintf("错误：%v, %v, %v", v, r[i], rv))
		}
	}
}
