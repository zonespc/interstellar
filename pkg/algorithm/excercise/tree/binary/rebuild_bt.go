package binary

import (
	"interstellar/pkg/algorithm/structure"
)

//RebuildByPreInOrder 根据前序和中序序列重建二叉树
func RebuildByPreInOrder(pre []int, plow int, phigh int, in []int, ilow int, ihigh int) *structure.BinaryTreeNode {
	if plow > phigh || ilow > ihigh || phigh > len(pre)-1 || ihigh > len(in)-1 {
		return nil
	}
	root := structure.NewBinaryTreeNode(pre[plow])
	j := ilow
	for j <= ihigh {
		if in[j] == pre[plow] {
			root.SetLeftNode(RebuildByPreInOrder(pre, plow+1, plow+j-ilow, in, ilow, j-1))
			root.SetRightNode(RebuildByPreInOrder(pre, plow+j-ilow+1, phigh, in, j+1, ihigh))
			break
		}
		j++
	}
	return root
}
