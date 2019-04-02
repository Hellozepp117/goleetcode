package main

import (
	"fmt"
	a "goleetcode/algorithm"
)

func main() {
	list := []float64{1, 2, 3}
	result := a.HeapSort(list)
	fmt.Println(result)
}
