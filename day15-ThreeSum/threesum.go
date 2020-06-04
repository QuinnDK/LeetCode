package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

func main() {

	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}

func threeSum(nums []int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	len := len(nums)
	result := [][]int{}
	for i := 0; i < len-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len - 1
		for {
			if j >= k {
				break
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum > 0 {
				k--
				continue
			}
			if sum < 0 {
				j++
				continue
			}
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				for {
					if j < k && nums[j] == nums[j+1] {
						j++
					} else {
						break
					}
				}
				for {
					if j < k && nums[k] == nums[k-1] {
						k--
					} else {
						break
					}
				}
				j++
				k--
			}
		}
	}
	return result
}
