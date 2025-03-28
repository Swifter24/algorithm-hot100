package main

import "fmt"

// 定义节点结构体
type Node struct {
	data int   // 数据域
	next *Node // 指针域，指向下一个节点
}

// 定义链表结构体
type LinkedList struct {
	head *Node // 头指针
}

// 插入新节点（头插法）
func (list *LinkedList) InsertFront(value int) {
	newNode := &Node{data: value}
	newNode.next = list.head
	list.head = newNode
}

// 插入新节点（尾插法）
func (list *LinkedList) InsertBack(value int) {
	newNode := &Node{data: value}
	if list.head == nil {
		list.head = newNode
		return
	}
	curr := list.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

// 删除节点
func (list *LinkedList) Delete(value int) {
	if list.head == nil {
		return
	}
	if list.head.data == value {
		list.head = list.head.next
		return
	}
	curr := list.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}
	if curr.next != nil {
		curr.next = curr.next.next
	}
}

// 查找节点
func (list *LinkedList) Search(value int) *Node {
	curr := list.head
	for curr != nil {
		if curr.data == value {
			return curr
		}
		curr = curr.next
	}
	return nil
}

// 打印链表
func (list *LinkedList) Print() {
	curr := list.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.next
	}
	fmt.Println("nil")
}

// 测试
func main() {
	list := &LinkedList{}

	list.InsertBack(10)
	list.InsertBack(20)
	list.InsertBack(30)
	list.InsertFront(5)
	list.Print() // 5 -> 10 -> 20 -> 30 -> nil

	list.Delete(20)
	list.Print() // 5 -> 10 -> 30 -> nil

	node := list.Search(10)
	if node != nil {
		fmt.Println("找到节点:", node.data)
	} else {
		fmt.Println("没找到")
	}
}
