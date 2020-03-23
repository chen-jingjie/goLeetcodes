package reverse_7

import "testing"

type example struct {
	value   int
	reverse int
}

var examples []example

func init() {
	examples = append(examples, example{123, 321})
	examples = append(examples, example{1596, 6951})
	examples = append(examples, example{0, 0})
	examples = append(examples, example{-12614, -41621})
	examples = append(examples, example{2, 2})
	examples = append(examples, example{-3, -3})
	examples = append(examples, example{75489, 98457})
	examples = append(examples, example{1896, 6981})
}

func TestReverse(t *testing.T) {
	for _, e := range examples {
		if e.reverse != Reverse(e.value) {
			t.Error("反转失败")
		}
	}
}
