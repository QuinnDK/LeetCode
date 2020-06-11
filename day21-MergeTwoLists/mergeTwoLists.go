package main

import "fmt"

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。


示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

*/

func main() {
	var head = new(ListNode)
	head.Val = 1
	var node1 = new(ListNode)
	node1.Val = 2
	//var node2 = new(ListNode)
	//node2.Val = 3
	//var node3 = new(ListNode)
	//node3.Val = 4
	var node2 = new(ListNode)
	node2.Val = 5

	head.Next = node1
	node1.Next = node2

	var head1 = new(ListNode)
	head1.Val = 1
	var node11 = new(ListNode)
	node11.Val = 3
	var node12 = new(ListNode)
	node12.Val = 4
	head1.Next = node11
	node11.Next = node12

	result := mergeTwoLists(head, head1)

	Shownode(result)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	result := prehead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}
	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}

func Shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}
