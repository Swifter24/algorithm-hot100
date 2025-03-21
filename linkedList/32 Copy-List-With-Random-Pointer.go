package main

//深拷贝

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	newHead := head.Next
	cur := head
	for ; cur.Next.Next != nil; cur = cur.Next {
		clone := cur.Next
		cur.Next = clone.Next
		clone.Next = clone.Next.Next
	}
	cur.Next = nil
	return newHead
}
