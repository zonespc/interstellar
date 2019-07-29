package practice

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	height := []int{1, 2, 4, 3}
	fmt.Println(MaxArea(height))
}

func TestMaxArea1(t *testing.T) {
	height := []int{1, 1}
	fmt.Println(MaxArea1(height))
}

func TestMaxAreaByDualPointer(t *testing.T) {
	height := []int{1, 2, 4, 3}
	fmt.Println(MaxAreaByDualPointer(height))
}
