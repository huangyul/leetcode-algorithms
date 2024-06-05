package totalfruit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAl(t *testing.T) {
	testCases := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "test 1",
			input:  []int{1, 2, 1, 3},
			output: 3,
		},
		{
			name:   "test 2",
			input:  []int{0, 1, 2, 2},
			output: 3,
		},
		{
			name:   "test 3",
			input:  []int{1, 2, 3, 2, 2},
			output: 4,
		},
		{
			name:   "test 4",
			input:  []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4},
			output: 5,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := totalFruit(tc.input)
			assert.Equal(t, tc.output, res)
		})
	}
}
