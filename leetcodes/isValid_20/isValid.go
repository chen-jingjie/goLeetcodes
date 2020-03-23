package isValid_20

func IsValid(s string) bool {
	lenStr := len(s)
	str := make([]int32, lenStr)
	l := 0
	parentheseMap := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	for i := 0; i < lenStr; i++ {
		b := string(s[i])
		_, ok := parentheseMap[b]
		if ok != false {
			str[l] = int32(i)
			l++
		} else if l != 0 && parentheseMap[string(s[str[l-1]])] == b {
			str[l-1] = 0
			l--
		} else {
			return false
		}
	}

	return l == 0
}
