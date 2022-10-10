package max_depth

import (
	"container/list"

	"github.com/guanguoyintao/galgorithm/tree"
)

//BFSMaxDepth BFS
func BFSMaxDepth(root *tree.BinaryTreeNode) int {
	// 根节点如果为0直接返回0
	if root == nil {
		return 0
	}
	// 创建一个队列
	lt := list.New()
	// 把根节点放入队列
	lt.PushBack(root)
	// 声明深度变量
	var depth int
	//判断队列不为nil
	for lt.Len() > 0 {
		ln := lt.Len()
		for i := 0; i < ln; i++ {
			node := lt.Remove(lt.Front()).(*tree.BinaryTreeNode)
			if node.Right != nil {
				lt.PushBack(node.Right)
			}
			if node.Left != nil {
				lt.PushBack(node.Left)
			}
		}
		depth++
	}

	return depth
}
