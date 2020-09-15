package main

import "sort"

/*
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/
var res [][]int

func subsetsWithDup(nums []int) [][]int {
	res = make([][]int, 0)
	sort.Ints(nums)
	dfs([]int{}, nums, 0)
	return res
}

func dfs(temp, nums []int, start int) {
	tmp := make([]int, len(temp))
	copy(tmp, temp)
	res = append(res, tmp)
	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] { // skip
			continue
		}
		temp = append(temp, nums[i])
		dfs(temp, nums, i+1)
		temp = temp[:len(temp)-1]
	}
}
