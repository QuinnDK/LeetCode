package main

import (
	"strconv"
	"strings"
)

/*

 */
func restoreIpAddresses(str string) []string {

	var (
		dfs func(start int, track []string, str string)
		ans []string
	)

	if len(str) > 12 || len(str) < 4 {
		return []string{}
	}

	dfs = func(start int, track []string, str string) {
		if start == 4 {
			// 如果已经遍历了4次，并且字符串也用完了，意味着拿到结果了
			if len(str) == 0 {
				ans = append(ans, strings.Join(track, "."))
			}
			return
		}
		// 跟当前字符串的长度，计算每一段IP的最低长度
		size := len(str)
		// 每个IP段最长只会有3个数字
		for i := 1; i <= 3; i++ {
			// 如果剩下的字符，比剩下允许的IP段长度还要大，肯定不对，说明当前段还得往后取
			trackLen := len(track)
			if size-i > 3*(4-trackLen) {
				if len(str[:i]) > 1 && str[:i][0] == '0' {
					break
				}
				continue
			}
			// 防止溢出
			if i > len(str) {
				break
			}
			num, e := strconv.Atoi(str[:i])
			if e != nil || num > 255 {
				break
			}
			track = append(track, str[:i])
			dfs(start+1, track, str[i:])
			track = track[:len(track)-1]
			// 如果当前数是0的话，那么不能给下次作为组合了，直接退出
			if num == 0 {
				break
			}
		}
	}

	dfs(0, []string{}, str)

	return ans
}
