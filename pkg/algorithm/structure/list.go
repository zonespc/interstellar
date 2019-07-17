package structure

import "fmt"

//SingleListNode 单向链表节点结构
type SingleListNode struct {
	val  int
	next *SingleListNode
}

//NewSingleListNode 新建一个节点
func NewSingleListNode(val int) *SingleListNode {
	node := &SingleListNode{
		val:  val,
		next: nil,
	}
	return node
}

//Value 获取链表节点值
func (s *SingleListNode) Value() int {
	return s.val
}

//SetValue 设置链表节点值
func (s *SingleListNode) SetValue(val int) {
	s.val = val
}

//NextNode 获取下一个节点
func (s *SingleListNode) NextNode() *SingleListNode {
	return s.next
}

//SetNextNode 设置下一个节点
func (s *SingleListNode) SetNextNode(next *SingleListNode) {
	s.next = next
}

//BuildSingleList 根据数组构建单向链表
func BuildSingleList(a []int) *SingleListNode {
	if a == nil {
		return nil
	}
	head := NewSingleListNode(a[0])
	p := head
	for i := 1; i < len(a); i++ {
		tNode := NewSingleListNode(a[i])
		p.SetNextNode(tNode)
		p = p.next
	}
	return head
}

//VisitNode 访问节点
func (s *SingleListNode) VisitNode() {
	fmt.Printf("%2d ", s.val)
}

//SeqTraversal 顺序遍历单向链表节点
func (s *SingleListNode) SeqTraversal() {
	p := s
	for p != nil {
		p.VisitNode()
		p = p.next
	}
	fmt.Printf("\n")
}

//InvertTraversal 倒序遍历单向链表
func (s *SingleListNode) InvertTraversal() {
	p := s
	InvertTraversalByRecursive(p)
	fmt.Printf("\n")
	InvertTraversalByStack(p)
	fmt.Printf("\n")
}

//InvertTraversalByRecursive 递归后序遍历节点
func InvertTraversalByRecursive(p *SingleListNode) {
	if p == nil {
		return
	}
	InvertTraversalByRecursive(p.next)
	p.VisitNode()
}

//InvertTraversalByStack 利用栈递归遍历节点
func InvertTraversalByStack(p *SingleListNode) {
	s := NewStack()
	for p != nil {
		s.Push(p)
		p = p.next
	}
	for s.IsEmpty() == false {
		p := s.Pop().(*SingleListNode)
		p.VisitNode()
	}
}
