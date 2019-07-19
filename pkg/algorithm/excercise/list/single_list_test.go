package list

import (
	"fmt"
	"interstellar/pkg/algorithm/structure"
	"testing"
)

func TestReverseSingleList(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	head := structure.BuildSingleList(a)
	head.SeqTraversal()
	imp := NewSingleListImp()
	head = imp.ReverseSingleList(head)
	head.SeqTraversal()
	pk := imp.KthLastNode(head, 5)
	fmt.Println(pk)

	b := []int{1, 2, 3, 4, 5}
	bp := structure.BuildSingleList(b)
	c := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	cp := structure.BuildSingleList(c)
	cp.NextNode().NextNode().NextNode().SetNextNode(bp.NextNode().NextNode())
	bp.SeqTraversal()
	cp.SeqTraversal()
	commn := imp.FindFirstCommonNode(bp, cp)
	fmt.Println(commn)
}
