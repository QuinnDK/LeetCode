package main

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

*/

import "fmt"

func main() {

	s := "abcdcab"
	//lengthOfLongestSubstring(s)
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)

}

func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	res := 0
	for i, j := 0, 0; i < len(s); i++ {
		m[s[i]]++
		//fmt.Println(m[s[i]])
		for m[s[i]] > 1 {
			m[s[j]]--
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
