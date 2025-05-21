package dynamicProgramming

//给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//示例 1：
//输入：s = "(()"
//输出：2
//解释：最长有效括号子串是 "()"
//
//示例 2：
//输入：s = ")()())"
//输出：4
//解释：最长有效括号子串是 "()()"
//
//示例 3：
//输入：s = ""
//输出：0

func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	ans := 0
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}
