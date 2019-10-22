package twoSum

import (
	"github.com/stretchr/testify/assert"
	"testing"

)

type twoSum struct {
	in []int
	target int
	out []int
}

var twoSums = []twoSum{
	twoSum {
		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1},
	},
	twoSum {
		[]int{3, 6, 8, 10},
		14,
		[]int{1, 2},
	},
	twoSum {
		[]int{3, 6, 8, 10},
		16,
		[]int{1, 3},
	},
}


func TestTwoSum(t *testing.T ) {
	for _, ts := range twoSums {
		assert.Equal(t, ts.out, TwoSum(ts.in, ts.target), "they should be equal.")
	}
}

