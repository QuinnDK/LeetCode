package main

import "fmt"

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
示例:
输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。
*/
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]byte][]string)
	for _, s := range strs {
		arr := getArr(s)
		m[arr] = append(m[arr], s)
	}
	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func getArr(s string) (arr [26]byte) {
	cnt := len(s)
	for i := 0; i < cnt; i++ {
		arr[s[i]-'a']++
	}
	return
}
