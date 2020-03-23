// 从网上抄的python版本的，还没理解透
package longestPalindrome_5

import (
	"math"
)

func LongestPalindrome(s string) string {
	a := []rune(s)
	newStr := "#"
	for _, r := range a {
		newStr += string(r) + "#"
	}

	lenNewStr := len(newStr)
	p := make([]int, lenNewStr)
	var maxLen, maxLenL, maxLenR int

	C, R, rad := 0, -1, 0
	for i := 0; i < lenNewStr; i++ {
		if i <= R {
			rad = int(math.Min(float64(p[2*C-i]), float64(R-i)))
		}

		for i+rad+1 < lenNewStr && i-rad-1 >= 0 && newStr[i+rad+1] == newStr[i-rad-1] {
			rad += 1
		}

		p[i] = rad

		if i+rad > R {
			R, C = i+rad, i
		}

		if rad > maxLen {
			maxLen = rad
			maxLenL = i - rad - 1
			maxLenR = i + rad + 1
		}
	}
	if (maxLenL-1)/2+1 <= (maxLenR-1)/2 {
		return s[(maxLenL-1)/2+1 : (maxLenR-1)/2]
	}
	return ""
}
