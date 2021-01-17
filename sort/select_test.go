package sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelectSort(t *testing.T) {
	sample := []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	result := []int{-3, -1, 1, 2, 3, 4, 5, 7, 8}
	require.Equal(t, result, SelectSort(sample))
}
