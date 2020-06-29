package main

import (
	"fmt"
	"strconv"
)

/*
给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。

注意：整数序列中的每一项将表示为一个字符串。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
第一项是数字 1

描述前一项，这个数是 1 即 “一个 1 ”，记作 11

描述前一项，这个数是 11 即 “两个 1 ” ，记作 21

描述前一项，这个数是 21 即 “一个 2 一个 1 ” ，记作 1211

描述前一项，这个数是 1211 即 “一个 1 一个 2 两个 1 ” ，记作 111221

*/
func main() {
	result := countAndSay(5)
	fmt.Println(result)
}
func countAndSay(n int) string {
	var r, s string
	var count int
	if n == 1 {
		return "1"
	}
	if n == 2 {
		return "11"
	}

	s = countAndSay(n - 1)
	//初始化字符计数器
	count = 1
	//1. 对S从左到右遍历
	for i := 1; i <= len(s)-1; i++ {
		//1.1 如果s[i] != s[i-1] 则把前面的字符打印
		if s[i] != s[i-1] {
			r += fmt.Sprintf("%s%s", strconv.Itoa(count), string(s[i-1]))
			count = 1
		} else {
			//1.2 如果相等则计数器+1
			count++
		}
		//1.2 如果到末尾了，则打印当前字符
		if i+1 == len(s) {
			r += fmt.Sprintf("%s%s", strconv.Itoa(count), string(s[i]))
			s = r
			return s
		}
	}
	return s
}
