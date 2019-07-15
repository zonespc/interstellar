package structure

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	queue := NewStack()
	for i := 0; i < len(a); i++ {
		queue.Push(a[i])
		fmt.Println("栈长度=", queue.Length())
	}
	for queue.IsEmpty() == false {
		q := queue.Pop()
		fmt.Println("栈 q =", q)
	}
}
