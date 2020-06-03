package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z
*/
func main() {
	s := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(s)
	fmt.Println(result)

}
func longestCommonPrefix(strs []string) string {
	strNumber := len(strs)
	//测试空切片
	if strNumber == 0 {
		return ""
	}

	//取得最小的字符串长度,即最大可能的公共前缀长度
	maxPreLength := len(strs[0])
	for x := 1; x < strNumber; x++ {
		strLength := len(strs[x])
		if strLength < maxPreLength {
			if strLength == 0 {
				//存在空字符串
				return ""
			}
			maxPreLength = strLength
		}
	}

	//获取相同部分的长度
	i := 0
I:
	for ; i < maxPreLength; i++ {
		for j := 1; j < strNumber; j++ {
			if strs[j][i] != strs[j-1][i] {
				break I
			}
		}
	}

	return strs[0][0:i]
}
