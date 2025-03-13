package main

import "fmt"

//示例 1：
//输入：nums = [100,4,200,1,3,2]
//输出：4
//解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

//示例 2：
//输入：nums = [0,3,7,2,5,8,4,6,0,1]
//输出：9

//示例 3：
//输入：nums = [1,0,1,2]
//输出：3

func longestConsecutive(nums []int) int {
	t := map[int]bool{}
	for _, num := range nums {
		t[num] = true
	}
	largestStreak := 0
	for k, _ := range t {
		if !t[k-1] {
			currentNum := k
			currentStreak := 1

			for t[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > largestStreak {
				largestStreak = currentStreak
			}
		}
	}
	return largestStreak
}

func main() {
	test := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(test))
}
