package list

import (
	"interstellar/pkg/algorithm/structure"
	"math"
)

//SingleListImp 单向链表操作实例
type SingleListImp struct{}

//NewSingleListImp 示例化
func NewSingleListImp() *SingleListImp {
	return &SingleListImp{}
}

//ReverseSingleList 翻转
func (imp *SingleListImp) ReverseSingleList(head *structure.SingleListNode) *structure.SingleListNode {
	if head == nil || head.NextNode() == nil {
		return head
	}
	p := head
	head = head.NextNode()
	p.SetNextNode(nil)
	for head.NextNode() != nil {
		tmp := head
		head = head.NextNode()
		tmp.SetNextNode(p)
		p = tmp
	}
	head.SetNextNode(p)
	return head
}

//KthLastNode 倒数第K个节点
func (imp *SingleListImp) KthLastNode(head *structure.SingleListNode, K int) *structure.SingleListNode {
	if K <= 0 || head == nil {
		return nil
	}
	p1 := head
	p2 := head
	for i := 0; i < K-1; i++ {
		if p2 == nil {
			break
		}
		p2 = p2.NextNode()
	}
	if p2 == nil {
		return nil
	}
	for p2.NextNode() != nil {
		p1 = p1.NextNode()
		p2 = p2.NextNode()
	}
	return p1
}

//FindFirstCommonNode 找出两个链表的第一个公共节点
func (imp *SingleListImp) FindFirstCommonNode(head1 *structure.SingleListNode, head2 *structure.SingleListNode) *structure.SingleListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	p1 := head1
	p2 := head2
	len1 := 0
	len2 := 0
	for p1 != nil {
		len1++
		p1 = p1.NextNode()
	}
	for p2 != nil {
		len2++
		p2 = p2.NextNode()
	}
	p1 = head1
	p2 = head2
	d := len1 - len2
	if len1 > len2 {
		for i := 0; i < int(math.Abs(float64(d))); i++ {
			p1 = p1.NextNode()
		}
	} else {
		for i := 0; i < int(math.Abs(float64(d))); i++ {
			p2 = p2.NextNode()
		}
	}
	for p1 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.NextNode()
		p2 = p2.NextNode()
	}
	return p1
}
