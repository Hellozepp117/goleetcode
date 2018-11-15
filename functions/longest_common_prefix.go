package functions

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	n := len(strs[0])
	for _, v := range strs {
		if n > len(v) {
			n = len(v)
		}
	}
	var i int
	for i = 0; i < n; i++ {
		ok := true
		for _, v := range strs {
			if v[i] != strs[0][i] {
				ok = false
			}
		}
		if !ok {
			break
		}
	}
	return strs[0][0:i]
}
