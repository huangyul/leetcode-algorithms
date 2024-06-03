package removezeros

// removeZeros 283.移动零
func removeZeros(nums []int) []int {
	slow, fast := 0, 0
	for fast <= len(nums)-1 {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			// 防止没发生0转移时
			if slow != fast {
				nums[fast] = 0
			}
			nums[fast] = 0
			slow++
		}
		fast++
	}
	return nums
}
