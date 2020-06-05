package main

import (
	"fmt"
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在唯一答案。

示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
*/
func main() {
	var target int
	fmt.Scan(&target)
	nums := []int{-1, 2, 1, -4}
	result := threeSumClosest(nums, target)
	fmt.Println(result)

}
func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	sort.Ints(nums) // 排序
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			tmp := nums[i] + nums[left] + nums[right]
			if tmp > target { // 偏大
				right--
			} else if tmp < target { // 偏小
				left++
			} else { // 一致
				return target
			}
			if distance(tmp, target) < distance(res, target) { // 距离更近则替换
				res = tmp
			}
		}
	}
	return res
}

func distance(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
