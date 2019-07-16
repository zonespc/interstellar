package binary

import (
	"interstellar/pkg/algorithm/structure"
	"testing"
)

func TestRebuildByPreIn(t *testing.T) {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	in := []int{4, 2, 5, 1, 6, 3, 7}
	root := RebuildByPreInOrder(pre, 0, len(pre)-1, in, 0, len(in)-1)
	structure.LevelTraversal(root)
}
