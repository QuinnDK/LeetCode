package main

import "fmt"

/*
给出一个无重叠的 ，按照区间起始端点排序的区间列表。
在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

示例 1：
输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]

*/
func insert(intervals [][]int, newInterval []int) [][]int {
	// 非法直接返回
	if len(newInterval) < 1 {
		return intervals
	}

	// 之前是空的，直接加入
	if len(intervals) < 1 {
		return [][]int{newInterval}
	}

	// 先找到左边必然要留下的
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		i++
	}
	left := i

	// 找到右边必然要留下的
	i = len(intervals) - 1
	for i >= 0 && intervals[i][0] > newInterval[1] {
		i--
	}
	right := i

	// 特殊情况处理，这种可能是中间没有交叉的区间，也可能是最左边，也可能是在最右边
	// 但都可以认为是要新添加一个区间
	if left > right {
		return append(intervals[:left], append([][]int{newInterval}, intervals[left:]...)...)
	}
	// 计算新区间
	if intervals[left][0] < newInterval[0] {
		newInterval[0] = intervals[left][0]
	}
	if intervals[right][1] > newInterval[1] {
		newInterval[1] = intervals[right][1]
	}
	// 填入新区间，最左要保留的+新区间+最右要保留的
	return append(intervals[:left], append([][]int{newInterval}, intervals[right+1:]...)...)
}

func main() {
	intervals := [][]int{{1, 3}, {5, 9}}
	interval := []int{2, 5}
	res := insert(intervals, interval)
	fmt.Println(res)
}
