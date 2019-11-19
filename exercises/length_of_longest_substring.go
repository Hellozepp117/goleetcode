package exercises

// 暴力法
func lengthOfLongestSubstringV1(s string) int {
	n := len(s)
	resp := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if unique(s, i, j) && resp < j-i {
				resp = j - i
			}
		}
	}
	return resp
}

// 移动窗格法
func unique(s string, begin, end int) bool {
	exist := struct{}{}
	hashSet := make(map[byte]struct{})
	for i := begin; i < end; i++ {
		ch := s[i]
		if _, ok := hashSet[ch]; ok {
			return false
		}
		hashSet[ch] = exist
	}
	return true
}

//反向map法
func lengthOfLongestSubstringV2(s string) int {
	n := len(s)
	exist := struct{}{}
	hashSet := make(map[byte]struct{})
	resp := 0
	var i, j int
	for i < n && j < n {
		if _, ok := hashSet[s[j]]; !ok {
			hashSet[s[j]] = exist
			j++
			if resp < j-i {
				resp = j - i
			}
		} else {
			delete(hashSet, s[i])
			i++
		}
	}
	return resp
}

func lengthOfLongestSubstringV3(s string) int {
	n := len(s)
	m := make(map[byte]int)
	resp := 0
	var i, j int
	for ; j < n; j++ {
		if _, ok := m[s[j]]; ok {
			if i < m[s[j]] {
				i = m[s[j]]
			}
		}
		if resp < j-i+1 {
			resp = j - i + 1
		}
		m[s[j]] = j + 1
	}
	return resp
}
