package binaryTree

//示例 1：
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//示例 2：
//输入：root = []
//输出：[]
//
//示例 3：
//输入：root = [1]
//输出：[1]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return ans
}
