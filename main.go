package main

import "fmt"
import "./algorithm"

func main() {
	list := []float64{0, 9, 2, 4, 8, 1, 3, 6, 5, 7, 4.6, 99.3, -11, -15, 0, 0, 7.5}
	result := algorithm.HeapSort(list)
	fmt.Println(result)
}
