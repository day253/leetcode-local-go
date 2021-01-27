package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type problem struct {
	Q string
	A string
}

func TestLongestPalindrome(t *testing.T) {
	list := []problem{
		{
			Q: "babad",
			A: "bab",
		},
		{
			Q: "cbbd",
			A: "bb",
		},
		{
			Q: "a",
			A: "a",
		},
		{
			Q: "ac",
			A: "a",
		},
	}
	for _, item := range list {
		assert.Equal(t, item.A, longestPalindrome(item.Q))
	}
}
