package dynamicProgramming

//给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
//
//注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
//示例 1：
//输入: s = "leetcode", wordDict = ["leet", "code"]
//输出: true
//解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
//
//示例 2：
//输入: s = "applepenapple", wordDict = ["apple", "pen"]
//输出: true
//解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//注意，你可以重复使用字典中的单词。
//
//示例 3：
//输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//输出: false

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool, len(wordDict))
	maxLen := 0
	for _, word := range wordDict {
		words[word] = true
		maxLen = max(maxLen, len(word))
	}

	n := len(s)
	f := make([]bool, n+1)
	f[0] = true
	for i := 1; i <= n; i++ {
		for j := i - 1; j >= max(i-maxLen, 0); j-- {
			if words[s[j:i]] && f[j] {
				f[i] = true
				break
			}
		}
	}
	return f[n]
}
