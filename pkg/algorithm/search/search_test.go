package search

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type Input struct {
		a     []int
		key   int
		ok    bool
		index int
	}
	inputs := []Input{
		Input{
			a:     []int{1, 2, 3, 2, 5},
			key:   0,
			ok:    false,
			index: -1,
		},
	}
	for _, in := range inputs {
		ok, index := BinarySearch(in.a, in.key)
		if ok != in.ok || index != in.index {
			fmt.Printf("Error output:%t %d\n", ok, index)
		} else {
			fmt.Printf("Pass output:%t %d\n", ok, index)
		}
	}

}
