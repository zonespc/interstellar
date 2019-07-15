package structure

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	a := []int{}
	for i := 1; i <= 15; i++ {
		a = append(a, i)
	}
	btNode := ConstructTree(a)
	LevelTraversal(btNode)
	fmt.Printf("\n")
	/*
		PreorderTraversal(btNode)
		fmt.Printf("\n")
		PreorderTraversalByCirculation(btNode)
		fmt.Printf("\n")
		InorderTraversal(btNode)
		fmt.Printf("\n")
		InorderTraversalByCirculation(btNode)
		fmt.Printf("\n")

		PostorederTraversal(btNode)
		fmt.Printf("\n")
		PostorederTraversalByCirculation(btNode)
		fmt.Printf("\n")
	*/
	fmt.Printf("边缘:\n")
	EdgeTraversal(btNode)
	fmt.Printf("\n")
	fmt.Printf("S型遍历:\n")
	STraversal(btNode)
	fmt.Printf("\n")

}
