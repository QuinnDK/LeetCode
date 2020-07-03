package main

import "fmt"

/*
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

示例 1:
输入: [1,2,0]
输出: 3

示例 2:
输入: [3,4,-1,1]
输出: 2

示例 3:
输入: [7,8,9,11,12]
输出: 1

1循环数组做以下操作。
2如果nums[i]为i+1的话，什么都不做，说明这个正数已经在正确的位置。
3如果nums[i]<=0或者>len(nums)的话，说明此值已经越界，不在考虑范围，内容清零。
4如果nums[i]在[1,len(nums)]的范围内(可能满足第二步，skip)，
取出此值val，内容清零。
5移动到val-1的数组位置，备份val-1位置的值，
设定val-1位置的值val，目的是让val的正数归位。根据备份的值，继续第五步。
终止条件为，某位置的值<=0或者>len(nums)或者等于所在index+1的值。
6循环数组，某位置为0，则返回位置index+1。如果全部正数归位，则len(nums)+1

*/
func main() {
	nums := []int{7, 8, 9, 11, 12}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		s := nums[i]
		if s > len(nums) || s <= 0 {
			nums[i] = 0
		} else if i+1 != s {
			t := s
			nums[i] = 0
			for {
				b := nums[t-1]
				nums[t-1] = t
				if b > len(nums) || b <= 0 || nums[b-1] == b {
					break
				}
				t = b
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
