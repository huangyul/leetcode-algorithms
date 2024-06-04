package minsubarraylen

// minSubArrayLen [209] 长度最小的子数组
//
// 注意题目是【大于或等于】都可以满足
func minSubArrayLen(nums []int, target int) int {
	s, f := 0, 0
	min := len(nums) + 1
	sum := 0
	for f < len(nums) {
		sum += nums[f]
		for sum >= target {
			l := f - s + 1
			if min > l {
				min = l
			}
			sum -= nums[s]
			s++
		}
		f++
	}

	if min == len(nums)+1 {
		return 0
	}

	return min
}
