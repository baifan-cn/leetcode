package LeetCodeArray

func removeElement(nums []int, val int) int {
	slowPoint := 0
	for fastPoint := 0; fastPoint < len(nums); fastPoint++ {
		if nums[fastPoint] != val {
			nums[slowPoint] = nums[fastPoint]
			slowPoint++
		}
	}
	return slowPoint
}
