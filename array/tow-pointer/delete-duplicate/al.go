package removeDuplicates

// removeDuplicates 26. 删除有序数组中的重复项移除元素
//
// 要求原地移除
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 1
	for fast <= len(nums)-1 {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1
}
