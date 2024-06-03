package backspacecompare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackSpaceCompare(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		t      string
		output bool
	}{
		{
			name:   "test 1",
			s:      "ab#c",
			t:      "ad#c",
			output: true,
		},
		{
			name:   "test 2",
			s:      "ab##",
			t:      "c#d#",
			output: true,
		},
		{
			name:   "test 3",
			s:      "ab#",
			t:      "c",
			output: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := backSpaceCompare(tc.s, tc.t)
			assert.Equal(t, tc.output, res)
		})
	}
}
