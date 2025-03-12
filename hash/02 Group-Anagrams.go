package main

import (
	"fmt"
	"sort"
)

//示例 1:
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

//示例 2:
//输入: strs = [""]
//输出: [[""]]

// 示例 3:
// 输入: strs = ["a"]
// 输出: [["a"]]
func groupAnagrams(strs []string) [][]string {
	map1 := make(map[string][]string)
	for _, str := range strs {
		byteStr := []byte(str)
		sort.Slice(byteStr, func(i, j int) bool {
			return byteStr[i] < byteStr[j]
		})
		str1 := string(byteStr)
		map1[str1] = append(map1[str1], str)
	}
	ans := make([][]string, 0, len(map1))
	for _, v := range map1 {
		ans = append(ans, v)
	}
	return ans
}
func main() {
	test := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	answer := groupAnagrams(test)
	fmt.Println(answer)
}
