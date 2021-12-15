package LeetCodeArray

/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。请必须使用时间复杂度为 O(log n) 的算法。
*/

func searchInsert(nums []int, target int) int {
	//// 空数组 or target小于等于数组第一个元素
	//if len(nums) == 0 || nums[0] >= target {
	//	return 0
	//	// target大于数组最后一个元素
	//} else if nums[len(nums)-1] < target {
	//	return len(nums)
	//	// target在数组中，使用二分法
	//} else {
	//	for k, v := range nums {
	//		if v == target {
	//			return k
	//		} else if v > target {
	//			return k
	//		}
	//	}
	//}
	//
	//return 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
