package isValid_20

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	examples := map[string]bool{
		"()[]{}": true,
		"(]":     false,
		"([)]":   false,
		"{[]}":   true,
		"":       true,
		"1{}":    false,
		"{1}1":   false,
	}

	for s, b := range examples {
		if IsValid(s) != b {
			fmt.Println(s, IsValid(s))
			t.Error("错误")
		}
	}
}
