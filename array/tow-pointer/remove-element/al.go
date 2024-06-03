package removeElement

// removeElement 27.移除元素
//
// 要求原地移除
//
// 定义快慢指针，当快指针所在位置不等于目标值时，就将数值放到慢指针的位置，同时慢指针往后移动一位；最后返回慢指针
func removeElement(nums []int, target int) int {
	slow, fast := 0, 0
	for fast <= len(nums)-1 {
		if nums[fast] != target {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
