package excercise

import (
	"fmt"
	"testing"
)

func TestFindDisorderArrayMid(t *testing.T) {
	a := []int{5, 4, 3, 2}
	fmt.Println(FindDisorderArrayMid(a))
}

func TestFindMidByHeap(t *testing.T) {
	a := []int{1, 5, 4, 0, 2}
	FindMidByHeap(a)
}
