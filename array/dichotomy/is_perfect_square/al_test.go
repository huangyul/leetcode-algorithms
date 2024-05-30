package isPerfectSquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPerfectSquare(t *testing.T) {
	testCases := []struct {
		name   string
		num    int
		output bool
	}{
		{
			name:   "test 1",
			num:    16,
			output: true,
		},
		{
			name:   "test 2",
			num:    14,
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isPerfectSquare(tc.num)
			assert.Equal(t, tc.output, res)
		})
	}
}
