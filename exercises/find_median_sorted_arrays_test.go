package exercises

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	n := findMedianSortedArraysV1([]int{1, 3}, []int{2})
	fmt.Println(n)
}
