package functions

import (
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	resp := fourSum([]int{-1, 0, 1, 2, -1, -4}, -1)
	fmt.Println(resp)
}
