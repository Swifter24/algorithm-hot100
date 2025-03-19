package main

//示例 1：
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//示例 2：
//输入：head = [1], n = 1
//输出：[]
//
//示例 3：
//输入：head = [1,2], n = 1
//输出：[1]

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	left, right := dummy, dummy
	for ; n > 0; n-- {
		right = right.Next
	}
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next

	return dummy.Next
}
