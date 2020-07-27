package main

import "fmt"

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

*/
//最优解，时间复杂度O(n),空间复杂度O(1)
//核心思想：当nums[i-1]大于零时，对后面的子数组和才会起增加的作用，故累加上；否则重新计数。

func maxSubArray(nums []int) int {

	l := len(nums)
	max := nums[0]

	for i := 1; i < l; i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func main() {
	num := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(num))
}
