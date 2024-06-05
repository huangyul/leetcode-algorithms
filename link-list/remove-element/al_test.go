package removeelement

import (
	"testing"

	linklist "github.com/huangyul/leetcode-algorithms/link-list"
	"github.com/stretchr/testify/assert"
)

func TestAl(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		val    int
		output []int
	}{
		{
			name:   "test 1",
			input:  []int{1, 2, 6, 3, 4, 5, 6},
			val:    6,
			output: []int{1, 2, 3, 4, 5},
		},
		{
			name:   "test 2",
			input:  []int{},
			val:    1,
			output: []int{},
		},
		{
			name:   "test 3",
			input:  []int{7, 7, 7, 7},
			val:    7,
			output: []int{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h := linklist.NewLinkedListFromSlice(tc.input)
			res := removeelement(h, tc.val)
			nres := linklist.ValuesFromLinkedList(res)
			assert.Equal(t, tc.output, nres)
		})
	}
}

// 输入：head = [1,2,6,3,4,5,6], val = 6
// 输出：[1,2,3,4,5]
// 示例 2：

// 输入：head = [], val = 1
// 输出：[]
// 示例 3：

// 输入：head = [7,7,7,7], val = 7
// 输出：[]
