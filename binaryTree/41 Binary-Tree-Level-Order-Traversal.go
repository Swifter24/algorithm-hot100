package binaryTree

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：[[3],[9,20],[15,7]]
//
// 示例 2：
// 输入：root = [1]
// 输出：[[1]]
//
// 示例 3：
// 输入：root = []
// 输出：[]
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		vals := make([]int, size)
		for i := range vals {
			node := queue[0]
			queue = queue[1:]
			vals[i] = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, vals)
	}
	return ans
}
