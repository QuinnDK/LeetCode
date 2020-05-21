package main

import "fmt"

func main() {

	var s string = "abcdc"
	longestPalindrome(s)
}
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	max := 0
	result := ""
	for i := range s {
		len1, str1 := expandAroundCenter(s, i, i) //奇数位字符串
		if len1 > max {
			max = len1
			result = str1
		}
		len2, str2 := expandAroundCenter(s, i, i+1) //偶数位字符串
		if len2 > max {
			max = len2
			result = str2
		}
	}
	fmt.Printf(result)
	return ""
}

//对输入的字符串做处理，回文从他的中心向两边展开，
func expandAroundCenter(s string, left, right int) (int, string) {
	L := left
	R := right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1, s[L+1 : R] //
}
