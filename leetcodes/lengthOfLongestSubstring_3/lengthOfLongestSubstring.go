package lengthOfLongestSubstring_3

func LengthOfLongestSubstring(s string) int {
	tmp := make(map[byte]int)
	l := 0

	var i, j int
	lenS := len(s)

	for j < lenS && i < lenS - l {
		v := s[j]
		k, ok := tmp[v]

		if ok {
			if x := j - i; x > l {
				l = x
			}
			for i < k  {
				delete(tmp, s[i])
				i++
			}
			i = k + 1
		}

		tmp[v] = j
		j++
	}

	if x := j - i; x > l {
		return x
	}

	return l
}
