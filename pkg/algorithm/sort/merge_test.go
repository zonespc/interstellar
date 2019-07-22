package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	a := []int{5, 4, 3, 2, 1}
	fmt.Println(a)
	MergeSort(a)
	fmt.Println(a)
}
