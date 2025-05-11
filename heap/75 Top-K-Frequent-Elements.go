package heap

//给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
//示例 1:
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//
//示例 2:
//输入: nums = [1], k = 1
//输出: [1]


func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	maxCnt := 0
	for _, num := range nums {
		cnt[num]++
		maxCnt = max(maxCnt, cnt[num])
	}

	buckets := make([][]int, maxCnt+1)
	for c,x := range cnt {
		buckets[x] = append(buckets[x], c)
	}
	ans := make([]int, 0, k)
	for i := maxCnt; i >= 0&&len(ans)<k ; i-- {
		ans = append(ans, buckets[i]...)
	}
	return ans
}