package main

import "fmt"

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	num := newList(nums)
	res := rotateRight(num, 2)

	for res != nil {
		fmt.Print(res.Val, " ")
		res = res.Next
	}
}

//数组转列表
func newList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0], Next: nil}
	cur := head
	for i := 1; i < n; i++ {
		newNode := &ListNode{Val: nums[i], Next: nil}
		cur.Next = newNode
		cur = newNode
	}
	return head
}
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	n, p := 1, head
	for p.Next != nil {
		p = p.Next
		n++
	}
	p.Next = head
	k %= n
	for i := 1; i <= n-k; i++ {
		p = p.Next
	}
	head, p.Next = p.Next, nil
	return head
}
