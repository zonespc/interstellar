package practice

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	a := []int{-2, 0, 0, 2, 2}
	fmt.Println(ThreeSum(a))
}

func TestThreeSumClosest(t *testing.T) {
	a := []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println(ThreeSumClosest(a, 82))
}
