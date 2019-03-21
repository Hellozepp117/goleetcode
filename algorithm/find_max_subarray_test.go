package algorithm

import (
	"fmt"
	"testing"
)

func TestFindMaxSubarray(t *testing.T) {
	list := []float64{0, -1, 2, 4, -2, 1, -3, 6, -5, 4}
	low, high, sum := FindMaxSubarray(list, 0, len(list)-1)
	fmt.Println(low, high, sum)
}
