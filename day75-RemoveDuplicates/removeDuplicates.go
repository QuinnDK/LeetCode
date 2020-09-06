package day75_RemoveDuplicates

/*
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定 nums = [1,1,1,2,2,3],

函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,1,2,3,3],

函数应返回新长度 length = 7, 并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。

你不需要考虑数组中超出新长度后面的元素。
*/

func removeDuplicates(nums []int) int {
	var count int //重复出现的次数
	for i := 0; i < len(nums); i++ {

		if i > 0 && nums[i] == nums[i-1] {
			count++
			if count >= 2 { //当count>=2时，进行删除
				nums = append(nums[:i], nums[i+1:]...) //对数组进行删除
				count--                                //次数-1
				i--                                    //指针不变，还是指向当前数组下标
			}

		} else {
			count = 0 //否则将次数返回0
		}
	}

	return len(nums) //返回结果
}

//双指针
func removeDuplicates1(nums []int) int {
	i, j := 1, 2
	for ; j < len(nums); j++ {
		if nums[j] != nums[i-1] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
