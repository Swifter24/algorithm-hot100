package backtrack

import "slices"

// 给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//
// 示例 1：
// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
//
// 示例 2：
// 输入：s = "a"
// 输出：[["a"]]
func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func partition(s string) [][]string {
	n := len(s)
	path := []string{}
	var ans [][]string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, slices.Clone(path))
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(s, i, j) {
				path = append(path, s[i:j+1])
				dfs(j + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}
