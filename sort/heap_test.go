package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	for _, problem := range problems() {
		HeapSort(problem.question)
		// assert.Equal(t, problem.answer, problem.question)
	}
}
