package generatematrix

// generatematrix 59.螺旋矩阵
//
// 定义好四个指针，再加一个代表方向的
func generatematrix(n int) [][]int {
	res := make([][]int, n)
	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1
	d := 0
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	for left <= right && top <= bottom {
		if d == 0 {
			for i := left; i <= right; i++ {
				res[top][i] = num
				num++
			}
			top++
		}
		if d == 1 {
			for i := top; i <= bottom; i++ {
				res[i][right] = num
				num++
			}
			right--
		}
		if d == 2 {
			for i := right; i >= left; i-- {
				res[bottom][i] = num
				num++
			}
			bottom--
		}
		if d == 3 {
			for i := bottom; i >= top; i-- {
				res[i][left] = num
				num++
			}
			left++
		}
		d++
		d = d % 4
	}

	return res
}
