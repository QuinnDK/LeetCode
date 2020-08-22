package main

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

*/
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

		//fmt.Println(i,v)   //输出中间参数便于理解
		k, ok := m[target-v]
		if ok {

			//fmt.Println(k,ok,m[target-v])

			result = append(result, k)
			result = append(result, i)
		}
		m[v] = i
	}
	return result
}
