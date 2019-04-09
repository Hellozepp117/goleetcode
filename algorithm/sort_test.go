package algorithm

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	list := []float64{2, 4, 3, 6, 5, 1}
	result := InsertSort(list)
	fmt.Println(result)
}

func TestBubbleSort(t *testing.T) {
	list := []float64{0, 9, 2, 4, 8, 1, 3, 6, 5, 7, 4.6, 99.3, -11, -15, 0, 0, 7.5}
	result := BubbleSort(list)
	fmt.Println(result)
}

func TestSelectSort(t *testing.T) {
	list := []float64{2, 4, 1, 3, 6, 5, 7, 9}
	result := SelectSort(list)
	fmt.Println(result)
}

func TestMergeSort(t *testing.T) {
	list := []float64{0, 9, 2, 4, 8, 1, 3, 6, 5, 7}
	result := MergeSort(list, 0, len(list)-1)
	fmt.Println(result)
}

func TestHeapSort(t *testing.T) {
	list := []float64{0, 9, 2, 4, 8, 1, 3, 6, 5, 7}
	result := HeapSort(list)
	fmt.Println(result)
}

func TestQuickSort(t *testing.T) {
	list := []float64{0, 9, 2, 4, 8, 1, 3, 6, 5, 7}
	result := QuickSort(list, 0, len(list)-1)
	fmt.Println(result)
}

func TestCountingSort(t *testing.T) {
	list := []int{9, 9, 7, 6, 3}
	result := CountingSort(list, 9)
	fmt.Println(result)
}

func TestRadixSort(t *testing.T) {
	list := []int{90, 19, 27, 1, 36, 30, 106}
	result := RadixSort(list, 3)
	fmt.Println(result)
}
