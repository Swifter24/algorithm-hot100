package stack

//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//
//示例 1:
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//
//示例 2：
//输入： heights = [2,4]
//输出： 4

func largestRectangleArea(heights []int) int {
	stack := []int{-1}
	var ans int
	heights = append(heights, -1)
	for right,h := range heights {
		for len(stack) > 1 && h <= heights[stack[len(stack)-1]] {
			i := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			ans = max(ans,heights[i]*(right-left-1))
		}
		stack = append(stack, right)
	}
	return ans
}