package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type question struct {
	nums   []int
	target int
	answer []int
}

func TestTwoSum(t *testing.T) {
	list := []question{
		{
			nums:   []int{3, 2, 4},
			target: 6,
			answer: []int{1, 2},
		},
		{
			nums:   []int{3, 2, 4},
			target: 5,
			answer: []int{0, 1},
		},
		{
			nums:   []int{0, 8, 7, 3, 3, 4, 2},
			target: 11,
			answer: []int{1, 3},
		},
		{
			nums:   []int{0, 1},
			target: 1,
			answer: []int{0, 1},
		},
		{
			nums:   []int{0, 3},
			target: 5,
			answer: nil,
		},
	}

	for _, question := range list {
		assert.Equal(t, question.answer, twoSum(question.nums, question.target))
	}
}
