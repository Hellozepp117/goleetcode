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

	}
	return list
}

func SelectSort(list []float64) []float64 {
	if len(list) < 2 {
		return list
	}
	for i := 0; i < len(list)-1; i++ {
		k := i
		v := list[i]
		for j := i + 1; j < len(list); j++ {
			if list[j] < v {
				v = list[j]
				k = j
			}
		}
		list[i], list[k] = list[k], list[i]
	}
	return list
}
