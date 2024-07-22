package no367

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAl(t *testing.T) {
	testCases := []struct {
		name   string
		input  int
		output bool
	}{
		{
			name:   "test 1",
			input:  16,
			output: true,
		},
		{
			name:   "test 2",
			input:  14,
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := isPerfectSquare(tc.input)
			assert.Equal(t, tc.output, res)
		})
	}
}
