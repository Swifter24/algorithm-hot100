package binaryTree

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
//
// 示例 1：
// 输入：root = [1,2,5,3,4,null,6]
// 输出：[1,null,2,null,3,null,4,null,5,null,6]
//
// 示例 2：
// 输入：root = []
// 输出：[]
//
// 示例 3：
// 输入：root = [0]
// 输出：[0]
func flatten(root *TreeNode) {
	dfs(root)
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	leftTail := dfs(root.Left)
	rightTail := dfs(root.Right)
	if leftTail != nil {
		leftTail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	if rightTail != nil {
		return rightTail
	}
	if leftTail != nil {
		return leftTail
	}
	return root
}
