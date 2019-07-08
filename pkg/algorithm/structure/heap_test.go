package structure

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeapInt(t *testing.T) {
	h := &HeapInt{}
	heap.Init(h)
	fmt.Println(*h)
	for i := 10; i > 0; i-- {
		heap.Push(h, i)
	}
	fmt.Println(*h)
	for len(*h) > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Printf("\n")
}
