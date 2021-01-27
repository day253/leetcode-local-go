package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type problem struct {
	Q int
	A int
}

func TestReverse(t *testing.T) {
	list := []problem{
		{
			Q: 123,
			A: 321,
		},
		{
			Q: -123,
			A: -321,
		},
		{
			Q: 120,
			A: 21,
		},
		{
			Q: 0,
			A: 0,
		},
		{
			Q: 1534236469,
			A: 0,
		},
		{
			Q: -2147483648,
			A: 0,
		},
		{
			Q: -147483412,
			A: -2143847412,
		},
	}
	for _, item := range list {
		assert.Equal(t, item.A, reverse(item.Q))
	}
}
