package sortedsquares

// sortedSquares [977] 有序数组的平方
//
// 因为是有序数组，所有平方后的最大值肯定在首位或者末尾；定义两个指针分别指向首尾，再定义一个i从末尾开始，找到首位最大值填到数组的最后
func sortedSquares(nums []int) []int {
	l := len(nums)
	res := make([]int, l)
	start, end := 0, l-1
	i := end
	for i >= 0 {
		sq := nums[start] * nums[start]
		eq := nums[end] * nums[end]
		if sq > eq {
			res[i] = sq
			start++
		} else {
			res[i] = eq
			end--
		}
		i--
	}
	return res
}
