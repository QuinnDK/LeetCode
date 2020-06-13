package main

import "fmt"

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

*/

func main() {
	nums := []int{1, 4, 5}
	nums1 := []int{1, 3, 4}
	nums2 := []int{2, 6}
	List := []*ListNode{newList(nums), newList(nums1), newList(nums2)}
	result := mergeKLists(List)
	fmt.Println(result)
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	} else if n == 1 {
		return lists[0]
	}
	// 分治法
	return merge(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

// 每次只合并两个链表
func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = merge(l1, l2.Next)
		return l2
	}
	l1.Next = merge(l1.Next, l2)
	return l1
}
