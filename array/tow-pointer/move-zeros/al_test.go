package removezeros

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveZeros(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output []int
	}{
		// {
		// 	name:   "test 1",
		// 	input:  []int{0, 1, 0, 3, 12},
		// 	output: []int{1, 3, 12, 0, 0},
		// },
		// {
		// 	name:   "test 2",
		// 	input:  []int{0},
		// 	output: []int{0},
		// },
		{
			name:   "test 3",
			input:  []int{1, 0},
			output: []int{1, 0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := removeZeros(tc.input)
			assert.Equal(t, tc.output, res)
		})
	}
}
