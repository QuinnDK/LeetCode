package main

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head.Next

	for fast != nil {
		if slow.Val == fast.Val {
			fast = fast.Next
		} else {
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}
	slow.Next = nil
	return head
}
