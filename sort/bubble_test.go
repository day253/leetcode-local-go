package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	for _, problem := range problems() {
		assert.Equal(t, problem.answer, BubbleSort(problem.question))
	}
}
