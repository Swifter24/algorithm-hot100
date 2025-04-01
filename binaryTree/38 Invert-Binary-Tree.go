package binaryTree

//示例 1：
//输入：root = [4,2,7,1,3,6,9]
//输出：[4,7,2,9,6,3,1]
//
//示例 2：
//输入：root = [2,1,3]
//输出：[2,3,1]
//
//示例 3：
//输入：root = []
//输出：[]

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
