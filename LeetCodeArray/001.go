package LeetCodeArray

//func twoSum(nums []int, target int) []int {
//	var res []int
//	for k, v := range nums[0 : len(nums)-1] {
//		for i, j := range nums[k+1 : len(nums)] {
//			if v+j == target {
//				res = append(res, k, k+1+i)
//			}
//		}
//	}
//	return res
//}
//
func TestTwoSum(nums []int, target int) []int {
	res := twoSum(nums, target)
	return res
}

func twoSum(nums []int, target int) []int {
	result := make(map[int]int, len(nums)/2)
	for k, v := range nums {
		/*查看元素在集合中是否存在 */
		if t, ok := result[target-v]; ok {
			return []int{t, k}
		}
		result[v] = k
	}
	return []int{0, 0}
}
