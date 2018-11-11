package functions

func reverse(x int) int {
	const INT_MAX = 2147483647
	const INT_MIN = -2147483648
	var rev int
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > INT_MAX/10 || (rev == INT_MAX/10 && pop > 7) {
			return 0
		}
		if rev < INT_MIN/10 || (rev == INT_MIN/10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}
