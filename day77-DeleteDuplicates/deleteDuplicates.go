package main

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	cur := prev.Next

	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}

		if prev.Next != cur {
			cur = cur.Next
			prev.Next = cur
			continue
		} else {
			prev = prev.Next
			cur = prev.Next
		}
	}

	return dummy.Next
}
