package removeElement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
		output int
	}{
		{
			name:   "test 1",
			nums:   []int{3, 2, 2, 3},
			target: 3,
			output: 2,
		},
		{
			name:   "test 2",
			nums:   []int{0, 1, 2, 2, 3, 0, 4, 2},
			target: 2,
			output: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := removeElement(tc.nums, tc.target)
			assert.Equal(t, tc.output, res)
		})
	}
}
