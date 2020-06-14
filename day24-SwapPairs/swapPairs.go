package main

import "fmt"

/*
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。



示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		h   = &ListNode{Val: 0, Next: head}
		p   = h
		cur = h.Next
	)

	for cur != nil && cur.Next != nil {
		p.Next = cur.Next
		cur.Next = cur.Next.Next
		p.Next.Next = cur

		p = cur
		cur = cur.Next
	}

	return h.Next
}

func main() {
	nums := []int{1, 2, 3, 4}
	head := newList(nums)
	result := swapPairs(head)
	fmt.Println(result)

}
