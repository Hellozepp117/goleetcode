package functions

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]
	for i := 0; i+2 < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { // skip same result
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			threeSum := nums[i] + nums[j] + nums[k]
			if toAbs(threeSum-target) < toAbs(closestSum-target) {
				closestSum = threeSum
			}
			if threeSum == target {
				return target
			} else if threeSum > target {
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				k--
			} else {
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				j++
			}
		}
	}
	return closestSum
}

func toAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
