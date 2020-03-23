package addDigits_258

import (
	"fmt"
	"testing"
)

func TestAddDigits(t *testing.T) {
	e := map[int]int{
		38:   2,
		0:    0,
		10:   1,
		156:  3,
		8465: 5,
	}

	for i, i2 := range e {
		if AddDigits(i) != i2 {
			t.Error("错误")
		}
	}
}

func ExampleAddDigits() {
	fmt.Println(AddDigits(38))
	fmt.Println(AddDigits(0))
	// Output:
	// 2
	// 0
}