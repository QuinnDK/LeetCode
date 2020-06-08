package main

import (
	"fmt"
	"sort"
)

/*

给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	result := fourSum(nums, target)
	fmt.Println(result)
}

func threeSum(nums []int, target int) [][]int {
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
			sum := nums[i] + nums[j] + nums[k] - target
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

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	var ret [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		fmt.Println(nums[i], target-nums[i])
		r := threeSum(nums[i+1:], target-nums[i])
		for j := 0; j < len(r); j++ {
			ret = append(ret, append(r[j], nums[i]))
		}
	}

	return ret
}
