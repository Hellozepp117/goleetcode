package exercises

func maxArea(height []int) int {
	n := len(height)
	var i, j, area int
	for j = n - 1; i < j; {
		var a int
		if height[i] > height[j] {
			a = height[j] * (j - i)
			j--
		} else {
			a = height[i] * (j - i)
			i++
		}
		if area < a {
			area = a
		}
	}
	return area
}
