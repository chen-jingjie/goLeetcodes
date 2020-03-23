package climbStairs_70

import "testing"

type example struct {
	value  int
	result int
}

var examples []example

func init() {
	examples = append(examples, example{2, 2})
	examples = append(examples, example{0, 0})
	examples = append(examples, example{3, 3})
	examples = append(examples, example{10, 89})
	examples = append(examples, example{30, 1346269})
	examples = append(examples, example{15, 987})
}

func TestClimbStairs(t *testing.T) {
	for _, e := range examples {
		r := ClimbStairs(e.value)
		if e.result != r {
			t.Error("错误")
		}
	}
}
