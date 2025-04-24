package backtrack

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//示例 1：
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//示例 2：
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//
//示例 3：
//输入：nums = [1]
//输出：[[1]]

func permute(nums []int) [][]int {
	n := len(nums)
	path := make([]int, n)
	onPath := make([]bool, n)
	var ans [][]int
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j, on := range onPath {
			if !on {
				path[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}
	dfs(0)
	return ans
}
