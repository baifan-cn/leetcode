package LeetCodeArray

import (
	"math"
	"sort"
)

/*
给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/

func MaxSubArray(nums []int) int {

	maxLocal, maxGlobal := 0.0, 0.0
	for _, v := range nums {
		maxLocal = math.Max(maxLocal+float64(v), float64(v))
		maxGlobal = math.Max(maxGlobal, maxLocal)
	}

	sort.Ints(nums)
	if nums[len(nums)-1] < 0 {
		return nums[len(nums)-1]
	}

	return int(maxGlobal)
}
