package longestCommonPrefix_14

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var prefix string

	for i := 0; true; i++ {
		var tmpStr byte
		for index, str := range strs {
			if i >= len(str) {
				return prefix
			}

			if index == 0 {
				tmpStr = str[i]
			} else if str[i] != tmpStr {
				return prefix
			}
		}
		prefix += string(tmpStr)
	}

	return prefix
}

// 上面在优化一版
//func LongestCommonPrefix(strs []string) string {
//	var prefix string
//
//	for _, str := range strs {
//		if str == "" { // 默认strs里面没有空字符串，否则就去掉这个注释
//			return ""
//		}
//		if prefix == "" {
//			prefix = str
//		} else {
//			var l int
//			var pre string
//			if len(prefix) < len(str) {
//				l = len(prefix)
//			} else {
//				l = len(str)
//			}
//			for i := 0; i < l; i++ {
//				if str[i] != prefix[i] {
//					break
//				}
//				pre += string(str[i])
//			}
//
//			if pre == "" {
//				return ""
//			} else {
//				prefix = pre
//			}
//		}
//	}
//
//	return prefix
//}
