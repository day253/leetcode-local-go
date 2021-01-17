package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type sortProblems struct {
	question []int
	answer   []int
}

func TestSort(t *testing.T) {
	problems := []sortProblems{
		{
			question: []int{3, 4, 5, 2, 1, 7, 8, -1, -3},
			answer:   []int{-3, -1, 1, 2, 3, 4, 5, 7, 8},
		},
		{
			question: []int{3, 4, 5, 2, 1},
			answer:   []int{1, 2, 3, 4, 5},
		},
	}
	for _, problem := range problems {
		assert.Equal(t, problem.answer, SelectSort(problem.question))
	}
}
