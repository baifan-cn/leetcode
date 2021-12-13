package main

import (
	"fmt"
	"github.com/baifan-cn/leetcode/LeetCodeArray"
)

func main() {
	nums := []int{2, 7, 11, 15}
	res := LeetCodeArray.TestTwoSum(nums, 9)
	fmt.Println(res)
}
