package sortedsquares

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "test 1",
			input:  []int{-4, -1, 0, 3, 10},
			output: []int{0, 1, 9, 16, 100},
		},
		{
			name:   "test 2",
			input:  []int{-7, -3, 2, 3, 11},
			output: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := sortedSquares(tc.input)
			assert.Equal(t, tc.output, res)
		})
	}
}
