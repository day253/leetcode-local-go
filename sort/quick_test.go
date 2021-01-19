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
