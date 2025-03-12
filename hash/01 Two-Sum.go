package main

import "fmt"

//示例 1：
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

//示例 2：
//输入：nums = [3,2,4], target = 6
//输出：[1,2]

//示例 3：
//输入：nums = [3,3], target = 6
//输出：[0,1]

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for index, value := range nums {
		if p, ok := hashmap[target-value]; ok {
			return []int{p, index}
		}
		hashmap[value] = index
	}
	return nil
}
func main() {
	answer := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(answer)
}
