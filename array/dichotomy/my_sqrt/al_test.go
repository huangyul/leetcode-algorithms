package mySqrt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mySqrt(t *testing.T) {
	testCases := []struct {
		name   string
		num    int
		output int
	}{
		{
			name:   "test 1",
			num:    4,
			output: 2,
		},
		{
			name:   "test 2",
			num:    8,
			output: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := mySqrt(tc.num)
			assert.Equal(t, tc.output, res)
		})
	}
}
