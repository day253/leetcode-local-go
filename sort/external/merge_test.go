package external

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	for _, problem := range problems() {
		assert.Equal(t, problem.answer, MergeSort(problem.question))
	}
}
