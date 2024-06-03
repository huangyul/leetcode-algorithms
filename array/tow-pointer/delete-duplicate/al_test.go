package removeDuplicates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		output int
	}{
		{
			name:   "test 1",
			nums:   []int{1, 1, 1, 2},
			output: 2,
		},
		{
			name:   "test 2",
			nums:   []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			output: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := removeDuplicates(tc.nums)
			assert.Equal(t, tc.output, res)
		})
	}
}
