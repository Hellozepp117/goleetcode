package exercises

// 暴力法
func longestPalindromeV1(s string) string {
	n := len(s)
	var resp string
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			ok := true
			p, q := i, j
			for p < q {
				if s[p] != s[q] {
					ok = false
					break
				}
				p++
				q--
			}
			if ok && len(resp) < j-i+1 {
				resp = string(s[i : j+1])
			}
		}
	}
	return resp
}
