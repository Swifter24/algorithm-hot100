package backtrack

import "slices"

//给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
//解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
//
//示例 1：
//输入：nums = [1,2,3]
//输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
//示例 2：
//输入：nums = [0]
//输出：[[],[0]]

func subsets(nums []int) [][]int {
	n := len(nums)
	path := make([]int, 0, n)
	ans := make([][]int, 0, 1<<n)
	var dfs func(int)

	dfs = func(i int) {
		ans = append(ans, slices.Clone(path))
		for j := i; i < n; j++ {
			path = append(path, nums[j])
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
