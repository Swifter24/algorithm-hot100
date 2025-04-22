package graph

//给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//此外，你可以假设该网格的四条边均被水包围。
//
//示例 1：
//输入：grid = [
//["1","1","1","1","0"],
//["1","1","0","1","0"],
//["1","1","0","0","0"],
//["0","0","0","0","0"]
//]
//输出：1
//
//示例 2：
//输入：grid = [
//["1","1","0","0","0"],
//["1","1","0","0","0"],
//["0","0","1","0","0"],
//["0","0","0","1","1"]
//]
//输出：3

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var ans int
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != '1' {
			return
		}
		grid[i][j] = '2'
		dfs(i, j-1)
		dfs(i, j+1)
		dfs(i-1, j)
		dfs(i+1, j)
	}
	for i, row := range grid {
		for j, col := range row {
			if col == '1' {
				dfs(i, j)
				ans++
			}
		}
	}
	return ans
}
