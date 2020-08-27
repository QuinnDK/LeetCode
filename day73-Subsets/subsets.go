package main

import "fmt"

/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

var ans [][]int

func subsets(nums []int) [][]int {
	ans = make([][]int, 0)
	curr := []int{}
	backtrack(nums, 0, curr)
	return ans
}

func backtrack(nums []int, start int, curr []int) {
	//if start == len(nums) {
	//	tmp := make([]int, len(curr))
	//	copy(tmp, curr)
	//	ans = append(ans, tmp)
	//	return
	//}
	//
	//backtrack(nums,  start+1, curr)
	//curr = append(curr, nums[start])
	//backtrack(nums,  start+1, curr)
	//curr = curr[:len(curr)-1]

	if start == len(nums) {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		ans = append(ans, tmp)
		return
	}
	//for i := start; i <len(nums); i++ {
	// 做选择
	backtrack(nums, start+1, curr)
	curr = append(curr, nums[start])
	// 回溯
	backtrack(nums, start+1, curr)

	curr = curr[:len(curr)-1]

}

//}
func main() {
	num := []int{1, 2, 3}
	fmt.Println(subsets(num))
}
