package no27

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAl(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
		val  int
		res  int
	}{
		// {
		// 	name: "test 1",
		// 	nums: []int{3, 2, 2, 3},
		// 	val:  3,
		// 	res:  2,
		// },
		// {
		// 	name: "test 2",
		// 	nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
		// 	val:  2,
		// 	res:  5,
		// },
		{
			name: "test 3",
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			res:  5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := removeElement(tc.nums, tc.val)
			assert.Equal(t, tc.res, res)
		})
	}
}
