package no69

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAl(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output int
	}{
		{
			name:   "test 1",
			input:  4,
			output: 2,
		},
		{
			name:   "test 3",
			input:  9,
			output: 3,
		},
		{
			name:   "test 2",
			input:  8,
			output: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := mySqrt(tc.input)
			assert.Equal(t, tc.output, res)
		})
	}
}
