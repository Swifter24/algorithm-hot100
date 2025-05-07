package binarySearch

//给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//如果数组中不存在目标值 target，返回 [-1, -1]。
//你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
//
//示例 1：
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//
//示例 2：
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//
//示例 3：
//输入：nums = [], target = 0
//输出：[-1,-1]

func searchRange(nums []int, target int) []int {
	if left(nums, target) == -1 {
		return []int{-1, -1}
	} else {
		return []int{left(nums, target), right(nums, target)}
	}
}

func left(nums []int, target int) int {
	l, r := 0, len(nums)-1
	candidate := -1
	for l <= r {
		mid := (r + l) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			candidate = mid
			r = mid - 1
		}
	}
	return candidate
}

func right(nums []int, target int) int {
	l, r := 0, len(nums)-1
	candidate := -1
	for l <= r {
		mid := (r + l) >> 1
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			candidate = mid
			l = mid + 1
		}
	}
	return candidate
}
