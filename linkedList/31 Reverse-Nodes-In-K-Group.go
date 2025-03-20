package main

//示例 1：
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//示例 2：
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	p0 := dummy
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	cur := p0.Next
	var pre *ListNode = nil
	for ; n >= k; n -= k {
		for i := 0; i < k; i++ {
			temp := cur.Next
			cur.Next = pre
			pre = cur
			cur = temp
		}
		temp := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = temp
	}
	return dummy.Next
}
