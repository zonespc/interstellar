package excercise

import (
	"fmt"
	"testing"
)

func TestFindKthMax(t *testing.T) {
	a := []int{10, 1, 2, 3, 0, 5}
	fmt.Println(FindKthMax(a, 6))
}
