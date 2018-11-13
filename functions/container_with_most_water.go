package functions

func maxArea(height []int) int {
	n := len(height)
	var i, j, area int
	for j = n - 1; i < j; {
		var a int
		if height[i] > height[j] {
			a = (height[i] - height[j]) * (j - i)
		} else {
			a = (height[j] - height[i]) * (j - i)
		}
		if area < a {
			area = a
		}
		if height[i+1] > height[j-1] {
			i++
		} else {
			j--
		}
	}
	return area
}
