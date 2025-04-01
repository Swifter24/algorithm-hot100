package binaryTree

//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//示例 1：
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//示例 2：
//输入：root = [1,2,2,null,3,null,3]
//输出：false

func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
}
func isSymmetric(root *TreeNode) bool {
	return isSameTree(root.Left, root.Right)
}
