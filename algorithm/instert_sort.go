package algorithm

// 插入排序
func InsertSort(list []float64) []float64 {
	if len(list) < 2 {
		return list
	}
	for i := 1; i < len(list); i++ {
		key := list[i]
		j := i - 1
		for ; j >= 0 && list[j] > key; j-- {
			list[j+1] = list[j]
		}
		list[j+1] = key
	}
	return list
}
