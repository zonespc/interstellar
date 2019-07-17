package structure

import "testing"

func TestSingleList(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	pHead := BuildSingleList(a)
	pHead.SeqTraversal()
	pHead.InvertTraversal()
}
