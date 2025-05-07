package binarySearch

//给你一个满足下述两条属性的 m x n 整数矩阵：
//每行中的整数从左到右按非严格递增顺序排列。
//每行的第一个整数大于前一行的最后一个整数。
//给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
//
//示例 1：
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//输出：true
//
//示例 2：
//输入：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//输出：false

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
