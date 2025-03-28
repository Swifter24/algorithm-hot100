package main
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//
//输入：head = []
//输出：[]


func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head2 := middle(head)
	head = sortList(head)
	head2 = sortList(head2)
	return merge(head, head2)
}
func middle(head *ListNode) *ListNode {
	pre, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next = nil
	return slow
}
func merge(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}
	if head1 != nil {
		cur.Next = head1
	} else {
		cur.Next = head2
	}
	return dummy.Next
}

///**
// * Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
// */
//// 876. 链表的中间结点（快慢指针）
//func middleNode(head *ListNode) *ListNode {
//	pre, slow, fast := head, head, head
//	for fast != nil && fast.Next != nil {
//		pre = slow // 记录 slow 的前一个节点
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//	pre.Next = nil  // 断开 slow 的前一个节点和 slow 的连接
//	return slow
//}
//
//// 21. 合并两个有序链表（双指针）
//func mergeTwoLists(list1, list2 *ListNode) *ListNode {
//	dummy := ListNode{} // 用哨兵节点简化代码逻辑
//	cur := &dummy // cur 指向新链表的末尾
//	for list1 != nil && list2 != nil {
//		if list1.Val < list2.Val {
//			cur.Next = list1 // 把 list1 加到新链表中
//			list1 = list1.Next
//		} else { // 注：相等的情况加哪个节点都是可以的
//			cur.Next = list2 // 把 list2 加到新链表中
//			list2 = list2.Next
//		}
//		cur = cur.Next
//	}
//	// 拼接剩余链表
//	if list1 != nil {
//		cur.Next = list1
//	} else {
//		cur.Next = list2
//	}
//	return dummy.Next
//}
//
//func sortList(head *ListNode) *ListNode {
//	// 如果链表为空或者只有一个节点，无需排序
//	if head == nil || head.Next == nil {
//		return head
//	}
//	// 找到中间节点 head2，并断开 head2 与其前一个节点的连接
//	// 比如 head=[4,2,1,3]，那么 middleNode 调用结束后 head=[4,2] head2=[1,3]
//	head2 := middleNode(head)
//	// 分治
//	head = sortList(head)
//	head2 = sortList(head2)
//	// 合并
//	return mergeTwoLists(head, head2)
//}
