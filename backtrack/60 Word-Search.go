package backtrack

import "slices"

//给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//示例 1：
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
//输出：true
//
//示例 2：
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
//输出：true
//
//示例 3：
//输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
//输出：false

var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func exist(board [][]byte, word string) bool {
	cnt := map[byte]int{}
	for _, row := range board {
		for _, c := range row {
			cnt[c]++
		}
	}

	w := []byte(word)
	wordCnt := map[byte]int{}
	for _, c := range w {
		wordCnt[c]++
		if wordCnt[c] > cnt[c] {
			return false
		}
	}

	if cnt[w[len(w)-1]] < cnt[w[0]] {
		slices.Reverse(w)
	}

	m, n := len(board), len(board[0])
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != w[k] {
			return false
		}
		if k == len(w)-1 {
			return true
		}
		board[i][j] = 0
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if x >= 0 && x < m && y >= 0 && y < n && dfs(x, y, k+1) {
				return true
			}
		}
		board[i][j] = w[k]
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
