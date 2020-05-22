package main

import "fmt"

func main() {

	var nums = []int{1, 3, 5, 7, 9}
	var target int = 8

	var results []int
	var results1 []int

	results = twoSum1(nums, target)

	results1 = twoSum2(nums, target)

	fmt.Println(results)
	fmt.Println(results1)

}

// 算法1：暴力破解法
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 算法2：map法
func twoSum2(nums []int, target int) []int {
	result := []int{}
	m := map[int]int{}
	for i, v := range nums {
		if k, ok := m[target-v]; ok {
			result = append(result, k)
			result = append(result, i)
		}
		m[v] = i
	}
	return result
}
