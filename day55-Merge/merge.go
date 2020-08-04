package main

import (
	"fmt"
	"sort"
)

/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	for i := 0; i < len(intervals); i++ {
		if len(result) == 0 || intervals[i][0] > result[len(result)-1][1] {
			result = append(result, intervals[i])
		} else if intervals[i][1] > result[len(result)-1][1] {
			result[len(result)-1][1] = intervals[i][1]
		}
	}
	return result
}

func main() {
	num := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(num))
}
