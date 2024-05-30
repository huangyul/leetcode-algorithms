package searchRange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		output []int
	}{
		{
			name:   "test 1",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			output: []int{3, 4},
		},
		{
			name:   "test 2",
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			output: []int{-1, -1},
		},
		{
			name:   "test 3",
			nums:   []int{},
			target: 0,
			output: []int{-1, -1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := searchRange(tc.nums, tc.target)
			assert.Equal(t, tc.output, res)
		})
	}
}
