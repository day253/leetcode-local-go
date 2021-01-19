package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectSort(t *testing.T) {
	for _, problem := range problems() {
		SelectSort(problem.question)
		assert.Equal(t, problem.answer, problem.question)
	}
}
