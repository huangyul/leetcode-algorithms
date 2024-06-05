package totalfruit

// totalFruit 904.水果成篮
//
// 使用窗口滑动算法，用两个指针定义滑动窗口前后，使用map记录种类及个数
func totalFruit(nums []int) int {
	maxLen := 0
	s, f := 0, 0
	count := make(map[int]int)

	for f <= len(nums)-1 {
		count[nums[f]]++
		// 如果超过两种水果，就缩小边界
		for len(count) > 2 {
			count[nums[s]]--
			if count[nums[s]] == 0 {
				delete(count, nums[s])
			}
			s++
		}
		if maxLen < (f - s + 1) {
			maxLen = f - s + 1
		}
		f++
	}

	return maxLen
}
