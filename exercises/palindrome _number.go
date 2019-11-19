package exercises

func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	const INT_MAX = 2147483647
	const INT_MIN = -2147483648
	var rev int
	n := x
	for n != 0 {
		pop := n % 10
		n /= 10
		if rev > INT_MAX/10 || (rev == INT_MAX/10 && pop > 7) {
			return false
		}
		if rev < INT_MIN/10 || (rev == INT_MIN/10 && pop < -8) {
			return false
		}
		rev = rev*10 + pop
	}
	if x == rev {
		return true
	}
	return false
}
