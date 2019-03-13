package algorithm

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	list := []float64{}
	result := InsertSort(list)
	fmt.Println(result)
}
