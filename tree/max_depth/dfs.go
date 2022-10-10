package max_depth

import "github.com/guanguoyintao/galgorithm/tree"

// DFSMaxDepth DFS
func DFSMaxDepth(root *tree.BinaryTreeNode) int {
	//如果root为nil，直接返回0
	if root == nil {
		return 0
	}
	//分别统计左节点深度、右节点深度，取最大值最后加上+1
	return max(DFSMaxDepth(root.Left), DFSMaxDepth(root.Right)) + 1
}
