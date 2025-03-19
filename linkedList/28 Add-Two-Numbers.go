package main

//示例 1：
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//示例 2：
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//示例 3：
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	node := 0
	for l1 != nil || l2 != nil || node != 0 {
		if l1 != nil {
			node += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			node += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: node % 10}
		node /= 10
		cur = cur.Next
	}
	return dummy.Next
}
