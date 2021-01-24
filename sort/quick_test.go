package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	for _, problem := range problems() {
		QuickSort(problem.question)
		assert.Equal(t, problem.answer, problem.question)
	}
}

func TestQuickSortRandSplit(t *testing.T) {
	for _, problem := range problems() {
		QuickSortRandSplit(problem.question)
		assert.Equal(t, problem.answer, problem.question)
	}
}

func TestThreeWayRadixQuickSort(t *testing.T) {
	for _, problem := range problems() {
		ThreeWayRadixQuickSort(problem.question)
		assert.Equal(t, problem.answer, problem.question)
	}
}
