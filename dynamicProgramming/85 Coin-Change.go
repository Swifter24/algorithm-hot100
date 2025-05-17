package dynamicProgramming

import "math"

//给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
//计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
//你可以认为每种硬币的数量是无限的。
//
//示例 1：
//输入：coins = [1, 2, 5], amount = 11
//输出：3
//解释：11 = 5 + 5 + 1
//
//示例 2：
//输入：coins = [2], amount = 3
//输出：-1
//
//示例 3：
//输入：coins = [1], amount = 0
//输出：0

func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	for i := range f {
		f[i] = math.MaxInt / 2
	}
	f[0] = 0
	for _, x := range coins {
		for c := x; c <= amount; c++ {
			f[c] = min(f[c], f[c-x]+1)
		}
	}
	ans := f[amount]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}
