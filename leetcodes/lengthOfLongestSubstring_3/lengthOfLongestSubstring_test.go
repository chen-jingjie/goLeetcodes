package lengthOfLongestSubstring_3

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	values := []string{"abcabcbb", "bbbbb", "pwwkew", "", "dvdf"}
	results := []int{3, 1, 3, 0, 3}

	for i, value := range values {
		if LengthOfLongestSubstring(value) != results[i] {
			fmt.Println(value, LengthOfLongestSubstring(value), results[i] )
			t.Error("错误")
		}
	}
}
