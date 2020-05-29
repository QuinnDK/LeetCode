package main

import (
	"fmt"
	"strconv"
)

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

*/
func main() {
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		return
	}
	Result := isPalindrome(i)
	fmt.Println(Result)
}

func isPalindrome(x int) bool {
	var str = strconv.Itoa(x)
	if x < 0 {
		return false
	}
	for i := range str {
		if str[len(str)-1-i] != str[i] {
			return false
		}
	}
	return true
}
