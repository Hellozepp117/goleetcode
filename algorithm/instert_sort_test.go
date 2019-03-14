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
	list := []float64{2, 1}
	result := BubbleSort(list)
	fmt.Println(result)
}
