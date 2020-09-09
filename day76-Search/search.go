package main

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,0,1,2,2,5,6] 可能变为 [2,5,6,0,0,1,2] )。

编写一个函数来判断给定的目标值是否存在于数组中。若存在返回 true，否则返回 false。

示例 1:

输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true
*/
func search(nums []int, target int) bool {
	if len(nums) <= 0 {
		return false
	}

	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}
		//1 无法判断那边递增，left++
		if nums[mid] == nums[left] {
			left++
		} else if nums[mid] > nums[left] { //2 说明左边递增
			if nums[mid] > target && nums[left] <= target { //2.1 通过固定两端可以肯定target，在此区间内
				right = mid - 1
			} else { //2.2 否则肯定在另一个区间内
				left = mid + 1
			}
		} else { //3 说明右区间递增
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
