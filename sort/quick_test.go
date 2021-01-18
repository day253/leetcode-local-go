package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	for _, problem := range problems() {
		assert.Equal(t, problem.answer, QuickSort(problem.question))
	}
}
