package main

import "fmt"
import "./algorithm"

func main() {
	list := []int{3, 3, 0, 2}
	result := algorithm.CountingSort(list, 3)
	fmt.Println(result)
}
