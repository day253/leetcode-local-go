package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type sortProblems struct {
	question [][]int
	answer   []int
}

func problems() []sortProblems {
	return []sortProblems{
		{
			question: [][]int{
				{1, 4, 5},
				{1, 3, 4},
			},
			answer: []int{1, 1, 3, 4, 4, 5},
		},
	}
}

func TestIntArray2List(t *testing.T) {
	arr := []int{1, 4, 5}
	assert.Equal(t, arr, List2IntArray(IntArray2List(arr)))
}

func TestMergeListRecursive(t *testing.T) {
	for _, problem := range problems() {
		linkListArr := []*ListNode{}
		for _, arr := range problem.question {
			linkListArr = append(linkListArr, IntArray2List(arr))
		}
		assert.Equal(t, problem.answer, List2IntArray(MergeListRecursive(linkListArr[0], linkListArr[1])))
	}
}

func TestMergeListIterate(t *testing.T) {
	for _, problem := range problems() {
		linkListArr := []*ListNode{}
		for _, arr := range problem.question {
			linkListArr = append(linkListArr, IntArray2List(arr))
		}
		assert.Equal(t, problem.answer, List2IntArray(MergeListIterate(linkListArr[0], linkListArr[1])))
	}
}
