package binaryTree

// 给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 小的元素（从 1 开始计数）。
// 示例 1：
// 输入：root = [3,1,4,null,2], k = 1
// 输出：1
//
// 示例 2：
// 输入：root = [5,3,6,2,4,null,null,1], k = 3
// 输出：3
func kthSmallest(root *TreeNode, k int) int {
	var ans int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		k--
		if k == 0 {
			ans = root.Val
		}
		dfs(root.Right)
	}
	dfs(root)
	return ans
}
