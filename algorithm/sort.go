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

func MergeSort(list []float64, start, end int) []float64 {
	if start >= end {
		return list
	}
	mid := (start + end) / 2
	list = MergeSort(list, start, mid)
	list = MergeSort(list, mid+1, end)
	list = mergeSublist(list, start, mid, end)
	return list
}

func mergeSublist(list []float64, start, mid, end int) []float64 {
	l1 := mid - start + 1
	l2 := end - mid
	left := make([]float64, l1)
	right := make([]float64, l2)
	for i := 0; i < l1; i++ {
		left[i] = list[start+i]
	}
	for i := 0; i < l2; i++ {
		right[i] = list[mid+1+i]
	}
	var i, j int
	for k := start; k <= end; k++ {
		if i == l1 {
			list[k] = right[j]
			j++
		} else if j == l2 {
			list[k] = left[i]
			i++
		} else if left[i] <= right[j] {
			list[k] = left[i]
			i++
		} else {
			list[k] = right[j]
			j++
		}
	}
	return list
}

//堆排序 heap-sort
func parent(i int) int {
	return i / 2
}
func left(i int) int {
	return 2 * i
}
func right(i int) int {
	return 2*i + 1
}

// 确保大顶子堆
func maxHeapify(list []float64, i, hSize int) []float64 {
	l := left(i)
	r := right(i)
	var largest int
	if l <= hSize {
		if list[l-1] > list[i-1] {
			largest = l
		}
	} else {
		largest = i
	}
	if r <= hSize {
		if list[r-1] > list[i-1] {
			largest = r
		}
		largest = r
	}
	if largest != i {
		list[i-1], list[largest-1] = list[largest-1], list[i-1]
		list = maxHeapify(list, largest, hSize)
	}
	return list
}

//建堆
func buildMaxHeap(list []float64) []float64 {
	hSize := len(list)
	for i := len(list); i >= 1; i-- {
		list = maxHeapify(list, i, hSize)
	}
	return list
}

func HeapSort(list []float64) []float64 {
	if len(list) == 1 {
		return list
	}
	hSize := len(list)
	list = buildMaxHeap(list)
	for i := len(list) / 2; i >= 2; i-- {
		list[0], list[i-1] = list[i-1], list[0]
		hSize--
		list = maxHeapify(list, 1, hSize)
	}
	return list
}
