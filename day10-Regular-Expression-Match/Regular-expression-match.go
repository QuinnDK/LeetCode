package main

import "fmt"

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

*/

func main() {
	s := "aa"
	p := "a."
	result := isMatch(s, p)
	fmt.Println(result)

}
func isMatch(s string, p string) bool {
	pLen := len(p)
	sLen := len(s)
	if pLen == 0 && sLen == 0 {
		return true
	}
	if pLen == 0 {
		return false
	}
	if p == ".*" {
		return true
	}
	return matchSwitchState(0, 0, s, p)
}

func matchSwitchState(si int, pi int, s string, p string) bool {
	pLen := len(p)
	sLen := len(s)
	//同时到达尾端 唯一可能返回true的情况
	if pi >= pLen && si >= sLen {
		return true
	}
	//匹配器没有了 但是字符还有
	if pi >= pLen && si < sLen {
		return false
	}
	//分情况向下递归 分别是 星号匹配 / 点号匹配 / 精确字符匹配
	if pi+1 < len(p) && p[pi+1] == byte('*') {
		mateByte := p[pi]
		//匹配超出 但星号匹配是可以忽略并继续步进到下一个匹配器的
		if si >= sLen {
			return matchSwitchState(si, pi+2, s, p)
		}
		//匹配到了有三种种选择 注意，匹配有两种情况 比较特殊的一种是'.*'可以和任意字符匹配
		//三种步进方式都有可能
		if s[si] == mateByte || mateByte == byte('.') {
			return matchSwitchState(si, pi+2, s, p) || matchSwitchState(si+1, pi, s, p) || matchSwitchState(si+1, pi+2, s, p)
		}
		//没有匹配到则只能匹配器继续步进下去，注意，*匹配器步进要向后两步
		return matchSwitchState(si, pi+2, s, p)
	} else if p[pi] == byte('.') {
		//匹配超出
		if si >= sLen {
			return false
		}
		//点号 可匹配任何字符 直接继续双步进即可
		return matchSwitchState(si+1, pi+1, s, p)
	} else {
		//匹配超出
		if si >= sLen {
			return false
		}
		//具体字符 只能严格匹配
		if p[pi] != s[si] {
			return false
		}
		return matchSwitchState(si+1, pi+1, s, p)
	}
}
