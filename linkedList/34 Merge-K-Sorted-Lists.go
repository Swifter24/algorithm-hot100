package main

import "container/heap"

//给你一个链表数组，每个链表都已经按升序排列。
//请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//示例 1：
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//示例 2：
//输入：lists = []
//输出：[]
//
//示例 3：
//输入：lists = [[]]
//输出：[]

func mergeKLists(lists []*ListNode) *ListNode {
	h := hp{}
	for _, list := range lists {
		if list != nil {
			h = append(h, list)
		}
	}
	heap.Init(&h)
	dummy := &ListNode{}
	cur := dummy
	for h.Len() > 0 {
		node := heap.Pop(&h).(*ListNode)
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

type hp []*ListNode

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(*ListNode)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
