package main

//示例 1：
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//示例 2：
//输入：head = []
//输出：[]
//
//示例 3：输入：head = [1]
//输出：[1]

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for head != nil && head.Next != nil {
		first, second := head, head.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
		head = first.Next
	}
	return dummy.Next
}
