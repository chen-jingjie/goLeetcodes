package isPalindrome_9

func IsPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	} else if x < 10 {
		return true
	}

	t := x
	y := 0
	for t != 0 {
		y = 10*y + t%10
		t /= 10
	}
	return x == y
}
