package no704

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
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			output: 4,
		},
		{
			name:   "test 2",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			output: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := search(tc.nums, tc.target)
			assert.Equal(t, tc.output, res)
		})
	}
}
