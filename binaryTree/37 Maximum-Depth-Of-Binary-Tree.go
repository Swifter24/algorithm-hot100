package binaryTree

//示例 1：
//输入：root = [3,9,20,null,null,15,7]
//输出：3
//
//示例 2：
//输入：root = [1,null,2]
//输出：2

func maxDepth(root *TreeNode) int {
	var ans int
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		depth++
		ans = max(ans, depth)
		dfs(root.Left, depth)
		dfs(root.Right, depth)
	}
	dfs(root, 0)
	return ans
}
