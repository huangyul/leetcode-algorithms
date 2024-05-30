package searchRange

// searchRange 34.在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			// 找到等于target的值后，分别处理左右指针，直到不等于target为止
			left, right = mid, mid
			for left >= 0 && nums[left] == target {
				left--
			}
			for right <= len(nums)-1 && nums[right] == target {
				right++
			}
			return []int{left + 1, right - 1}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1, -1}
}
