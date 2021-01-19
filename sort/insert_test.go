package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	for _, problem := range problems() {
		InsertSort(problem.question)
		assert.Equal(t, problem.answer, problem.question)
	}
}
