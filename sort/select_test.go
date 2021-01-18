package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectSort(t *testing.T) {
	for _, problem := range problems() {
		assert.Equal(t, problem.answer, SelectSort(problem.question))
	}
}
