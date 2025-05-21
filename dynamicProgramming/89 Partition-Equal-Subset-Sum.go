package dynamicProgramming

//给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。
//
//示例 1：
//输入：nums = [1,5,11,5]
//输出：true
//解释：数组可以分割成 [1, 5, 5] 和 [11] 。
//
//示例 2：
//输入：nums = [1,2,3,5]
//输出：false
//解释：数组不能分割成两个元素和相等的子集。

func canPartition(nums []int) bool {
	s := 0
	for _, n := range nums {
		s += n
	}
	if s%2 != 0 {
		return false
	}
	s /= 2
	s2 := 0
	f := make([]bool, s+1)
	f[0] = true

	for _, i := range nums {
		s2 = min(s2+i, s)
		for j := s2; j >= i; j-- {
			f[j] = f[j] || f[j-i]
		}
		if f[s] {
			return true
		}
	}
	return false
}
