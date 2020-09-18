package main

/*
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:
输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n || head.Next == nil {
		return head
	}
	count := 0         // 计数到第几个链表
	var newhead = head // 记录链表开头
	var pre *ListNode
	var headpoint1, headpoint2 = head, head.Next // 记录反转开始时的两个节点

	for head != nil {
		count++
		// 如果遍历到需要反转的链表段
		if count >= m && count <= n {
			// 如果是最后一个需要反转的节点
			if count == n {
				if m == 1 {
					newhead.Next = head.Next
					newhead = head
				} else {
					headpoint2.Next = head.Next
					headpoint1.Next = head
				}
			}
			// 反转链表
			tem := head.Next
			head.Next = pre
			pre = head
			head = tem
		} else {
			headpoint1 = head
			headpoint2 = head.Next

			head = head.Next // 继续遍历
		}

	}
	return newhead
}

//* 1. 头插法反转链表, 每次将遍历的当前节点插入开始反转的位置。
//* 2. 通过哨兵节点处理 m == 1 的情况
//* 3. 通过m的值确定头插法的头节点位置
//* 4. 通过n-m的值确定执行几次头插操作

func reverseBetween1(head *ListNode, m int, n int) *ListNode {
	dummy, i, j := &ListNode{Next: head}, m, n-m
	d := dummy
	for i > 1 {
		d = d.Next
		i--
	}
	cur := d.Next.Next
	pre := d.Next
	for j > 0 {
		pre.Next = cur.Next
		cur.Next = d.Next
		d.Next = cur
		cur = pre.Next
		j--
	}
	return dummy.Next
}
