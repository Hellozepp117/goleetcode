package exercises

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i+2 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { // skip same result
			continue
		}
		j := i + 1
		k := len(nums) - 1
		target := -nums[i]
		for j < k {
			if nums[j]+nums[k] > target {
				k--
			} else if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] { // skip same result
					j++
				}
				for j < k && nums[k] == nums[k+1] { // skip same result
					k--
				}
			} else {
				j++
			}
		}
	}
	return ans
}
