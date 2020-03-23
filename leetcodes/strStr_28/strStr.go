package strStr_28

func StrStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	l1 := len(haystack)
	l2 := len(needle)
	for i := 0; i <= l1-l2; i++ {
		v := true
		for i2 := 0; i2 < l2; i2++ {
			if needle[i2] != haystack[i+i2] {
				v = false
				break
			}
		}
		if v {
			return i
		}
	}
	return -1
}
