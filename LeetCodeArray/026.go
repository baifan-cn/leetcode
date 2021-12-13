package LeetCodeArray

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[count] != nums[i] {
			count += 1
			nums[count] = nums[i]
		}
	}
	fmt.Println(nums[0 : count+1])
	return count + 1
}

func TestRemoveDuplicates(nums []int) {
	n := removeDuplicates(nums)
	fmt.Printf("n=%d", n)
}
