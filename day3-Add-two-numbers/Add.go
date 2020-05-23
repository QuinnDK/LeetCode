package main

import "fmt"

/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := newList([]int{8, 9, 9}) // 5	// 9 8
	l2 := newList([]int{2})       // 5	// 1
	cur := addTwoNumbers(l1, l2)
	for cur != nil {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
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

//链表相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 定义一个 i 表示 两数相加是否大于10, 定义一个s表示两数两加之和
	var i, s int
	// 定义一个头节点
	res := &ListNode{Val: 0}
	// 定义一个当前节点
	now := res

	// 无限循环,内部加条件跳出
	for true {
		// 如果 i >0 表示前面的数之和大于10了,所以当前的和要加1,否则不加
		if i > 0 {
			s = l1.Val + l2.Val + 1
		} else {
			s = l1.Val + l2.Val
		}
		// 如果 两数之和大于10,那么该位的值应该是 s - 10 ,否则就是 和 本身,并设置 i 标记
		if s >= 10 {
			now.Next = &ListNode{Val: s - 10}
			i = 1
		} else {
			now.Next = &ListNode{Val: s}
			i = 0
		}
		// 将 当前节点移动到下一指针
		now = now.Next

		// 当l1 和 l2 都移动到最后了,准备跳出循环
		if l1.Next == nil && l2.Next == nil {
			// 如果l1 和 l2 最后的和大于10,即 i == 1, 那么后面还需要添加一个 1 (小学数学问题,不用解释为什么是1吧)
			if i == 1 {
				now.Next = &ListNode{Val: 1}
			}
			// 跳出循环
			break
		}
		// 如果 l1走到了最后,可能l2 还没结束,所以把l1的当前节点值设置成0,继续跟l2相加,否则移动到下一指针
		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}
		// 如果 l2走到了最后,可能l1 还没结束,所以把l2的当前节点值设置成0,继续跟l1相加,否则移动到下一指针
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}
	}

	// 返回头节点的下一个节点指针,因为头节点是我设置成0的
	return res.Next
}
