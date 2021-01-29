package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type problem struct {
	Q string
	A int
}

func TestMyAtoi(t *testing.T) {
	list := []problem{
		{
			Q: "42",
			A: 42,
		},
		{
			Q: "   -42",
			A: -42,
		},
		{
			Q: "4193 with words",
			A: 4193,
		},
		{
			Q: "words and 987",
			A: 0,
		},
		{
			Q: "-91283472332",
			A: -2147483648,
		},
		{
			Q: "-2147483647",
			A: -2147483647,
		},
	}
	for _, item := range list {
		assert.Equal(t, item.A, myAtoi(item.Q))
	}
}
