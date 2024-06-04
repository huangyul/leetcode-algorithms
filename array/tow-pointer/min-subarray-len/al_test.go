package minsubarraylen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAl(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		output int
	}{
		{
			name:   "test 1",
			nums:   []int{2, 3, 1, 2, 4, 3},
			target: 7,
			output: 2,
		},
		{
			name:   "test 2",
			nums:   []int{1, 4, 4},
			target: 4,
			output: 1,
		},
		{
			name:   "test 3",
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			target: 11,
			output: 0,
		},
		{
			name:   "test 4",
			nums:   []int{1, 2, 3, 4, 5},
			target: 11,
			output: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := minSubArrayLen(tc.nums, tc.target)
			assert.Equal(t, tc.output, res)
		})
	}
}
