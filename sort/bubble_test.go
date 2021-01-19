package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	for _, problem := range problems() {
		BubbleSort(problem.question)
		assert.Equal(t, problem.answer, problem.question)
	}
}
