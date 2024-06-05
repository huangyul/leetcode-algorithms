package totalfruit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAl(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "test 1",
			input:  []int{1, 2, 1, 3},
			output: 3,
		},
		{
			name:   "test 2",
			input:  []int{0, 1, 2, 2},
			output: 3,
		},
		{
			name:   "test 3",
			input:  []int{1, 2, 3, 2, 2},
			output: 4,
		},
		{
			name:   "test 4",
			input:  []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			output: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := totalFruit(tc.input)
			assert.Equal(t, tc.output, res)
		})
	}
}

// 示例 1：

// 输入：fruits = [1,2,1]
// 输出：3
// 解释：可以采摘全部 3 棵树。
// 示例 2：

// 输入：fruits = [0,1,2,2]
// 输出：3
// 解释：可以采摘 [1,2,2] 这三棵树。
// 如果从第一棵树开始采摘，则只能采摘 [0,1] 这两棵树。
// 示例 3：

// 输入：fruits = [1,2,3,2,2]
// 输出：4
// 解释：可以采摘 [2,3,2,2] 这四棵树。
// 如果从第一棵树开始采摘，则只能采摘 [1,2] 这两棵树。
// 示例 4：

// 输入：fruits = [3,3,3,1,2,1,1,2,3,3,4]
// 输出：5
// 解释：可以采摘 [1,2,1,1,2] 这五棵树。
