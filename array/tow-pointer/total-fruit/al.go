package totalfruit

// totalFruit 904.水果成篮
//
// 使用窗口滑动算法，用两个指针定义滑动窗口前后，使用map记录种类及个数
func totalFruit(nums []int) int {
	maxLen := 0
	s, f := 0, 0
	m := make(map[int]int)
	for f < len(nums) {
		m[nums[f]]++
		for len(m) > 2 {
			m[nums[s]]--
			if m[nums[s]] == 0 {
				delete(m, nums[s])
			}
			s++
		}
		if (f - s + 1) > maxLen {
			maxLen = (f - s + 1)
		}
		f++
	}

	return maxLen

}
