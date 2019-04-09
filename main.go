package main

import "fmt"
import "./algorithm"

func main() {
	list := []int{90, 19, 27, 1, 36, 30, 106}
	result := algorithm.RadixSort(list, 3)
	fmt.Println(result)
}
