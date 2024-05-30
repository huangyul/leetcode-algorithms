package isPerfectSquare

// isPerfectSquare 367.完全平方数
func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		mid := (left + right) / 2
		x := mid * mid
		if x == num {
			return true
		} else if x > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
