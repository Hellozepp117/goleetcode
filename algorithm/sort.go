package algorithm

import "math"

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
		if ok == true {
			break
		}
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

// 确保大顶子堆(list 为大顶堆)
func maxHeapify(list []float64, i, hSize int) ([]float64, int) {
	l := left(i)
	r := right(i)
	var largest int
	if l <= hSize && list[l-1] > list[i-1] {
		largest = l
	} else {
		largest = i
	}
	if r <= hSize && list[r-1] > list[largest-1] {
		largest = r
	}
	if largest != i {
		list[i-1], list[largest-1] = list[largest-1], list[i-1]
		list, _ = maxHeapify(list, largest, hSize)
	}
	return list, largest
}

//建堆
func buildMaxHeap(list []float64) []float64 {
	hSize := len(list)
	n := len(list)
	for i := n/2 + 1; i >= 1; i-- {
		list, _ = maxHeapify(list, i, hSize)
	}
	return list
}

//堆排序 heap-sort
func HeapSort(list []float64) []float64 {
	if len(list) == 1 {
		return list
	}
	hSize := len(list)
	list = buildMaxHeap(list)
	for i := len(list); i >= 2; i-- {
		list[0], list[i-1] = list[i-1], list[0]
		hSize--
		list, _ = maxHeapify(list, 1, hSize)
	}
	return list
}

// 快速排序 quick-sort
func QuickSort(list []float64, p, r int) []float64 {
	if p < r {
		list, q := partition(list, p, r)
		list = QuickSort(list, p, q-1)
		list = QuickSort(list, q, r)
	}
	return list
}

func partition(list []float64, p, r int) ([]float64, int) {
	x := list[r]
	i := p - 1
	for j := p; j < r; j++ {
		if list[j] <= x {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[r] = list[r], list[i+1]
	return list, i + 1
}

// 计数排序(元素必为整数)
func CountingSort(list []int, k int) []int {
	temp := make([]int, k+1)
	for i := 0; i < len(list); i++ {
		temp[list[i]]++
		// temp[i] shows # of elements == i
	}
	for i := 1; i <= k; i++ {
		temp[i] += temp[i-1]
		// temp[i] shows # of elements <= i
	}
	output := make([]int, len(list))
	for j := len(list) - 1; j >= 0; j-- {
		output[temp[list[j]]-1] = list[j]
		temp[list[j]]--
	}
	return output
}

// 十进制基数排序(LSD)
func RadixSort(list []int, n int) []int {
	for i := 0; i < n; i++ {
		list = fillAndMergeBucks(list, i)
	}
	return list
}

func fillAndMergeBucks(list []int, n int) []int {
	bucks := make([][]int, 10)
	for i := 0; i < len(list); i++ {
		k := (list[i] / int(math.Pow10(n))) % 10
		bucks[k] = append(bucks[k], list[i])
	}
	var output []int
	for i := 0; i < 10; i++ {
		for _, v := range bucks[i] {
			output = append(output, v)
		}
	}
	return output
}
