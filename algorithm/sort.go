package algorithm

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

func BubbleSort(list []float64) []float64 {
	if len(list) < 2 {
		return list
	}
	for {
		ok := true
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
				ok = false
			}
		}
		if ok {
			break
		}
	}
	return list
}
