package main

import "fmt"

/*
一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。

示例 1:
输入: "12"
输出: 2
解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。

*/

func numDecodings(s string) int {
	l := len(s)
	if l == 0 {
		panic("题目说了非空字符串了")
	}
	dp := make([]int, l)
	for i := 0; i < l; i++ {
		dp[i] = 0
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				dp[i] = dp[i-1]
			} else {
				return 0
			}
		}
		if s[i-1] == '1' {
			dp[i] = dp[i-1] + dp[i-2]
		}
		if s[i-1] == '2' && '1' <= s[i] && s[i] <= '6' {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[l-1]
}
func main() {
	fmt.Println(numDecodings("123"))
}
