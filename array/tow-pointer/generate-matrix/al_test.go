package generatematrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAl(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output [][]int
	}{
		{
			name:  "test 1",
			input: 3,
			output: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name:  "test 2",
			input: 1,
			output: [][]int{
				{1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := generatematrix(tc.input)
			assert.Equal(t, tc.output, res)
		})
	}
}
