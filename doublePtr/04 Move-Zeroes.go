package main

import "fmt"

//示例 1:
//输入: nums = [0,1,0,3,12]
//输出: [1,3,12,0,0]

//示例 2:
//输入: nums = [0]
//输出: [0]

func MoveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	test := []int{0, 1, 0, 3, 12}
	MoveZeroes(test)
	fmt.Println(test)
}
