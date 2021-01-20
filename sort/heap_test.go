package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	for _, problem := range problems() {
		HeapSort(problem.question)
		assert.Equal(t, problem.answer, problem.question)
	}
}
