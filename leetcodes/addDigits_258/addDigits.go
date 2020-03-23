package addDigits_258

func AddDigits(num int) int {
	r := 0

	for num != 0 {
		v := num % 10
		num /= 10
		r += v
		if r >= 10 {
			r -= 9
		}
	}
	return r
}
