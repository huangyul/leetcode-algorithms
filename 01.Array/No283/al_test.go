package no283

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
			nums: []int{0, 1, 0, 3, 12},
			res:  []int{1, 3, 12, 0, 0},
		},
		{
			name: "test 2",
			nums: []int{0},
			res:  []int{0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := moveZeroes(tc.nums)
			assert.Equal(t, tc.res, res)
		})
	}
}
