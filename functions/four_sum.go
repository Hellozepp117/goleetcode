package functions

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for l := 0; l+3 < len(nums); l++ {
		if l > 0 && nums[l] == nums[l-1] {
			continue
		}
		newTarget := target - nums[l]
		for i := l + 1; i+2 < len(nums); i++ {
			if i > 0 && nums[i] == nums[i-1] { // skip same result
				continue
			}
			j := i + 1
			k := len(nums) - 1
			newTarget2 := newTarget - nums[i]
			for j < k {
				if nums[j]+nums[k] > newTarget2 {
					k--
				} else if nums[j]+nums[k] == newTarget2 {
					ans = append(ans, []int{nums[l], nums[i], nums[j], nums[k]})
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
	}
	return ans
}
