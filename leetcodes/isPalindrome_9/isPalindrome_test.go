package isPalindrome_9

import "testing"

func TestIsPalindrome(t *testing.T) {
	v := map[int]bool{
		121: true,
		-121: false,
		10: false,
		0: true,
		5: true,
		48464867646: false,
		12345678987654321: true,
	}

	for k, b := range v {
		if IsPalindrome(k) != b {
			t.Error("错误")
		}
	}
}