package exercises

// 暴力法
func findMedianSortedArraysV1(nums1 []int, nums2 []int) float64 {
	array := append(nums1, nums2...)
	n := len(array)
	for i := 0; i < n-1; i++ {
		done := true
		for j := 0; j < n-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				done = false
			}
		}
		if done {
			break
		}
	}
	var resp float64
	if n%2 == 1 {
		resp = float64(array[(n-1)/2])
	} else {
		resp = float64(array[n/2]+array[n/2-1]) / 2
	}
	return resp
}
