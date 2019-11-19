package exercises

// map法 36ms
func romanToInt(s string) int {
	romInt := make(map[byte]int)
	romInt['I'] = 1
	romInt['V'] = 5
	romInt['X'] = 10
	romInt['L'] = 50
	romInt['C'] = 100
	romInt['D'] = 500
	romInt['M'] = 1000
	var ans int
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romInt[s[i+1]] > romInt[s[i]] {
			ans -= romInt[s[i]]
		} else {
			ans += romInt[s[i]]
		}
	}
	return ans
}

// 暴力法 20ms
func romanToIntV2(s string) int {
	var ans, i int
	n := len(s)
	for {
		if i < n {
			if s[i] != 'M' {
				break
			} else {
				ans += 1000
				i++
			}
		} else {
			break
		}
	}
	if i < n-1 && s[i] == 'C' && s[i+1] == 'M' {
		ans += 900
		i += 2
	}
	if i < n && s[i] == 'D' {
		ans += 500
		i++
	}
	if i < n-1 && s[i] == 'C' && s[i+1] == 'D' {
		ans += 400
		i += 2
	}

	for {
		if i < n {
			if s[i] != 'C' {
				break
			} else {
				ans += 100
				i++
			}
		} else {
			break
		}
	}
	if i < n-1 && s[i] == 'X' && s[i+1] == 'C' {
		ans += 90
		i += 2
	}
	if i < n && s[i] == 'L' {
		ans += 50
		i++
	}
	if i < n-1 && s[i] == 'X' && s[i+1] == 'L' {
		ans += 40
		i += 2
	}

	for {
		if i < n {
			if s[i] != 'X' {
				break
			} else {
				ans += 10
				i++
			}
		} else {
			break
		}
	}
	if i < n-1 && s[i] == 'I' && s[i+1] == 'X' {
		ans += 9
		i += 2
	}
	if i < n && s[i] == 'V' {
		ans += 5
		i++
	}
	if i < n-1 && s[i] == 'I' && s[i+1] == 'V' {
		ans += 4
		i += 2
	}
	for {
		if i < n {
			if s[i] != 'I' {
				break
			} else {
				ans++
				i++
			}
		} else {
			break
		}
	}
	return ans
}
