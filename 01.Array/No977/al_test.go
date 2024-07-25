package no977

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAl(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "test 1",
			nums: []int{-4, -1, 0, 3, 10},
			res:  []int{0, 1, 9, 16, 100},
		},
		{
			name: "test 1",
			nums: []int{-7, -3, 2, 3, 11},
			res:  []int{4, 9, 9, 49, 121},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := sortedSquares(tc.nums)
			assert.Equal(t, tc.res, res)
		})
	}
}
