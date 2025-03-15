package main

import "fmt"

//示例 1：
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
//示例 2：
//输入：height = [1,1]
//输出：1

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		maxArea := min(height[l], height[r]) * (r - l)
		ans = max(ans, maxArea)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func main() {
	test := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(test))
}
