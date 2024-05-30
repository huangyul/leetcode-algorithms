package mySqrt

// mySqrt 69.x的平方根
func mySqrt(num int) int {
	// 0和1的平方根都是本身
	if num < 2 {
		return num
	}
	left, right := 2, num/2
	for left <= right {
		mid := (left + right) / 2
		x := mid * mid
		if x == num {
			return mid
		} else if x > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
