package no844

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAl(t *testing.T) {
	testCases := []struct {
		name string
		s    string
		t    string
		res  bool
	}{
		{
			name: "test 1",
			s:    "ab#c",
			t:    "ad#c",
			res:  true,
		},
		{
			name: "test 2",
			s:    "ab##",
			t:    "c#d#",
			res:  true,
		},
		{
			name: "test 3",
			s:    "a#c",
			t:    "b",
			res:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := backspaceCompare(tc.s, tc.t)
			assert.Equal(t, tc.res, res)
		})
	}
}
