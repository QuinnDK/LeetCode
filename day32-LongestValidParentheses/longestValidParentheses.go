package main

import "fmt"

/*
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

*/

func longestValidParentheses(s string) int {
	stack := []int{}
	stack = append(stack, -1)
	max := 0
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1] // 弹出栈顶元素
			if len(stack) == 0 {
				stack = append(stack, i)
			}
			// 当前元素的索引与栈顶元素作差，获取最近的括号匹配数
			max = Max(max, i-stack[len(stack)-1])
		}
	}
	return max
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	s := ")()())"
	result := longestValidParentheses(s)
	fmt.Println(result)
}
