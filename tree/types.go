package tree

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewPIBinaryTreeNode(preorder, inorder []int) *BinaryTreeNode {
	if len(preorder) < 1 {
		return nil
	}
	nodeVal := preorder[0]
	if len(preorder) == 1 {
		return &BinaryTreeNode{Val: nodeVal}
	}
	root := &BinaryTreeNode{Val: nodeVal}
	var index int

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == nodeVal {
			index = i
		}
	}
	preLeft, preRight := preorder[1:index+1], preorder[index+1:]
	inLeft, inRight := inorder[:index], inorder[index+1:]

	root.Left = NewPIBinaryTreeNode(preLeft, inLeft)
	root.Right = NewPIBinaryTreeNode(preRight, inRight)

	return root
}
