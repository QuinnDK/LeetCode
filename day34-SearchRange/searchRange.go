package main

import "fmt"

/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/
func searchRange(nums []int, target int) []int {
	left := -1
	right := -1

	if len(nums) == 0 || nums[len(nums)-1] < target {
		return []int{left, right}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			left = i
		}
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] == target {
			right = j
		}
	}
	return []int{right, left}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	targe := 8
	result := BinarySearch(nums, targe)
	fmt.Println(result)
}
func BinarySearch(nums []int, target int) []int {
	low := 0
	height := len(nums) - 1

	if len(nums) <= 0 {
		return []int{-1, -1}
	}

	//左右同时开始搜索，并且左边索引要小于右边索引
	for nums[low] != nums[height] && low < height {
		if nums[low] != target {
			if target > nums[low] {
				low = low + 1
			}
		}

		if nums[height] != target {
			if target < nums[height] {
				height = height - 1
			}
		}
	}

	if nums[low] == target && nums[height] == target {
		return []int{low, height}
	} else {
		return []int{-1, -1}
	}
}
