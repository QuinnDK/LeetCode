package main

import (
	"fmt"
	"sort"
)

/*
给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用一次。

说明：

所有数字（包括目标数）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
*/
func main() {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	res := combinationSum2(nums, 8)
	fmt.Println(res)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) //快排，懒得写
	res := [][]int{}
	dfs(candidates, nil, target, 0, &res) //深度优先
	return res
}

func dfs(candidates, nums []int, target, left int, res *[][]int) {
	if target == 0 { //结算
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for i := left; i < len(candidates); i++ { // left限定，形成分支
		if i != left && candidates[i] == candidates[i-1] { // *同层节点 数值相同，剪枝
			continue
		}
		if target < candidates[i] { //剪枝
			return
		}

		dfs(candidates, append(nums, candidates[i]), target-candidates[i], i+1, res) //*分支 i+1避免重复
	}
}
