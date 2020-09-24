package main

/*
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	st(root, &ret)
	return ret
}

func st(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	st(root.Left, ret)
	*ret = append(*ret, root.Val)
	st(root.Right, ret)
}
