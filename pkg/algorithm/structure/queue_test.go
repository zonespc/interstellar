package structure

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	queue := NewQueue()
	for i := 0; i < len(a); i++ {
		queue.Enqueue(a[i])
		fmt.Println("队列长度=", queue.Length())
	}
	for queue.IsEmpty() == false {
		q := queue.Dequeue()
		fmt.Println("队列 q =", q)
	}
}
