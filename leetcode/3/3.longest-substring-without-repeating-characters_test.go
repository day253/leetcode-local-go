package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type problem struct {
	Q string
	A int
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, item := range []problem{
		{
			Q: "abcabcbb",
			A: 3,
		}, {
			Q: "bbbbb",
			A: 1,
		}, {
			Q: "pwwkew",
			A: 3,
		}, {
			Q: "",
			A: 0,
		},
	} {
		assert.Equal(t, item.A, lengthOfLongestSubstring(item.Q))
	}
}
