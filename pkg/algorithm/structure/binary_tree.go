package structure

import (
	"fmt"
	"math"
)

//BinaryTreeNode 二叉树节点
type BinaryTreeNode struct {
	val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

//NewBinaryTreeNode 新建一个二叉树节点
func NewBinaryTreeNode(val int) *BinaryTreeNode {
	btNode := &BinaryTreeNode{
		val:   val,
		left:  nil,
		right: nil,
	}
	return btNode
}

//ConstructTree 根据数组按层次构建二叉树
func ConstructTree(a []int) *BinaryTreeNode {
	if a == nil {
		return nil
	}
	q := NewQueue()
	btNode := NewBinaryTreeNode(a[0])
	q.Enqueue(btNode)
	for i := 1; i < len(a); {
		tq := q.Dequeue().(*BinaryTreeNode)
		tq.left = NewBinaryTreeNode(a[i])
		q.Enqueue(tq.left)
		i++
		if i < len(a) {
			tq.right = NewBinaryTreeNode(a[i])
			q.Enqueue(tq.right)
			i++
		}
	}
	return btNode
}

//LevelTraversal 按层次遍历
func LevelTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	q := NewQueue()
	q.Enqueue(root)
	cnt := 0
	level := 0
	for q.IsEmpty() == false {
		tq := q.Dequeue().(*BinaryTreeNode)
		cnt++
		l := int(math.Floor(math.Log2(float64(cnt))))
		if l > level {
			fmt.Printf("\n")
			level = l
		}
		fmt.Printf("%2d ", tq.val)
		if tq.left != nil {
			q.Enqueue(tq.left)
		}
		if tq.right != nil {
			q.Enqueue(tq.right)
		}
	}
}

//PreorderTraversal 前序遍历
func PreorderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	fmt.Printf("%2d ", root.val)
	if root.left != nil {
		PreorderTraversal(root.left)
	}
	if root.right != nil {
		PreorderTraversal(root.right)
	}
}

//PreorderTraversalByCirculation 非递归前序遍历
func PreorderTraversalByCirculation(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	p := root
	q := NewStack()
	q.Push(p)
	for q.IsEmpty() == false {
		VisitNode(p)
		if p.right != nil {
			q.Push(p.right)
		}
		if p.left != nil {
			p = p.left
		} else {
			p = q.Pop().(*BinaryTreeNode)
		}
	}

}

//InorderTraversal 中序遍历
func InorderTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	if root.left != nil {
		InorderTraversal(root.left)
	}
	fmt.Printf("%2d ", root.val)
	if root.right != nil {
		InorderTraversal(root.right)
	}
}

//InorderTraversalByCirculation 非递归中序遍历
func InorderTraversalByCirculation(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	q := NewStack()
	p := root
	for q.IsEmpty() == false || p != nil {
		if p != nil {
			q.Push(p)
			p = p.left
		} else {
			p = q.Pop().(*BinaryTreeNode)
			fmt.Printf("%2d ", p.val)
			p = p.right
		}
	}
}

//PostorederTraversal 后序遍历
func PostorederTraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	if root.left != nil {
		PostorederTraversal(root.left)
	}
	if root.right != nil {
		PostorederTraversal(root.right)
	}
	fmt.Printf("%2d ", root.val)
}

//PostorederTraversalByCirculation 后序非递归调用
func PostorederTraversalByCirculation(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	q := NewStack()
	p := root
	for p != nil {
		q.Push(p)
		p = p.left

	}
	var pLastVisit *BinaryTreeNode
	for q.IsEmpty() == false {
		p = q.Pop().(*BinaryTreeNode)
		if p.right == nil || p.right == pLastVisit {
			fmt.Printf("%2d ", p.val)
			pLastVisit = p
		} else {
			q.Push(p)
			p = p.right
			for p != nil {
				q.Push(p)
				p = p.left
			}
		}
	}
}

//VisitNode 访问节点
func VisitNode(node *BinaryTreeNode) {
	if node != nil {
		fmt.Printf("%d ", node.val)
	}
}

//EdgeTraversal 边缘遍历
func EdgeTraversal(root *BinaryTreeNode) {
	edgeTraversalLeft(root)
	edgeTraversalBottom(root)
	edgeTraversalRight(root)
}
func edgeTraversalLeft(root *BinaryTreeNode) {
	p := root
	for p.left != nil {
		VisitNode(p)
		p = p.left
	}
}
func edgeTraversalBottom(root *BinaryTreeNode) {
	p := root
	q := NewStack()
	q.Push(p)
	for q.IsEmpty() == false {
		if p.left == nil && p.right == nil {
			VisitNode(p)
		}
		if p.right != nil {
			q.Push(p.right)
		}
		if p.left != nil {
			p = p.left
		} else {
			p = q.Pop().(*BinaryTreeNode)
		}
	}
}
func edgeTraversalRight(root *BinaryTreeNode) {
	if root == nil || root.right == nil {
		return
	}
	p := root.right
	q := NewStack()
	for p != nil {
		q.Push(p)
		p = p.right
	}
	q.Pop()
	for q.IsEmpty() == false {
		p = q.Pop().(*BinaryTreeNode)
		VisitNode(p)
	}
}

//STraversal S型遍历
func STraversal(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	s1 := NewStack()
	s2 := NewStack()
	p := root
	s1.Push(p)
	cnt := 1
	level := 0
	for s1.IsEmpty() == false || s2.IsEmpty() == false {
		if level%2 == 0 {
			p := s1.Pop().(*BinaryTreeNode)
			VisitNode(p)
			if p.right != nil {
				s2.Push(p.right)
			}
			if p.left != nil {
				s2.Push(p.left)
			}
		} else {
			p := s2.Pop().(*BinaryTreeNode)
			VisitNode(p)
			if p.left != nil {
				s1.Push(p.left)
			}
			if p.right != nil {
				s1.Push(p.right)
			}
		}
		cnt++
		level = int(math.Floor(math.Log2(float64(cnt))))
	}
}
