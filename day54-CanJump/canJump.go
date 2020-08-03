package main

import "fmt"

/*
给定一个非负整数数组，你最初位于数组的第一个位置。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个位置。

示例 1:

输入: [2,3,1,1,4]
输出: true
解释: 我们可以先跳 1 步，从位置 0 到达 位置 1, 然后再从位置 1 跳 3 步到达最后一个位置。
*/
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true

	for i := 1; i < len(nums); i++ {
		flag := false
		for j := 0; j < i; j++ {
			if dp[j] && nums[j]+j >= i {
				flag = true
				break
			}
		}
		dp[i] = flag
	}

	return dp[len(nums)-1]
}
func main() {
	num := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(num), canJump1(num))
}

func canJump1(nums []int) bool {
	l := len(nums) - 1
	for i := l - 1; i >= 0; i-- {
		if nums[i]+i >= l {
			l = i
		}
	}
	return l <= 0
}
