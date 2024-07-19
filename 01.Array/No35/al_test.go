package no35

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
			nums:   []int{1, 3, 5, 6},
			target: 5,
			output: 2,
		},
		{
			name:   "test 2",
			nums:   []int{1, 3, 5, 6},
			target: 2,
			output: 1,
		},
		{
			name:   "test 3",
			nums:   []int{1, 3, 5, 6},
			target: 7,
			output: 4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := searchInsert(tc.nums, tc.target)
			assert.Equal(t, tc.output, res)
		})
	}
}
