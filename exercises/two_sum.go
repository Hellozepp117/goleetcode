package exercises

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok && v != i {
			return []int{i, v}
		}
	}
	return nil
}
