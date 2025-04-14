package binaryTree

//给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
//示例 1：
//输入：root = [1,2,3,null,5,null,4]
//输出：[1,3,4]
//
//示例 2：
//输入：root = [1,2,3,4,null,null,null,5]
//输出：[1,3,4,5]
//
//示例 3：
//输入：root = [1,null,3]
//输出：[1,3]
//
//示例 4：
//输入：root = []
//输出：[]

func rightSideView(root *TreeNode) []int {
	var ans []int
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == len(ans) {
			ans = append(ans, root.Val)
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}
	dfs(root, 0)
	return ans
}
