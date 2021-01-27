package leetcode

import "math"

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (25.87%)
 * Likes:    4253
 * Dislikes: 6570
 * Total Accepted:    1.4M
 * Total Submissions: 5.3M
 * Testcase Example:  '123'
 *
 * Given a signed 32-bit integer x, return x with its digits reversed. If
 * reversing x causes the value to go outside the signed 32-bit integer range
 * [-2^31, 2^31 - 1], then return 0.
 *
 * Assume the environment does not allow you to store 64-bit integers (signed
 * or unsigned).
 *
 *
 * Example 1:
 * Input: x = 123
 * Output: 321
 * Example 2:
 * Input: x = -123
 * Output: -321
 * Example 3:
 * Input: x = 120
 * Output: 21
 * Example 4:
 * Input: x = 0
 * Output: 0
 *
 *
 * Constraints:
 *
 *
 * -2^31 <= x <= 2^31 - 1
 *
 *
 */

/*
https://golang.org/pkg/math/#pkg-constants
MaxInt8   = 1<<7 - 1 = 127
MinInt8   = -1 << 7 = -128
MaxInt16  = 1<<15 - 1 = 32767
MinInt16  = -1 << 15 = -32768
MaxInt32  = 1<<31 - 1 = 2147483647
MinInt32  = -1 << 31 = -2147483648
MaxInt64  = 1<<63 - 1
MinInt64  = -1 << 63
MaxUint8  = 1<<8 - 1 = 255
MaxUint16 = 1<<16 - 1 = 65535
MaxUint32 = 1<<32 - 1 = 4294967295
MaxUint64 = 1<<64 - 1
*/
/*
(ans * 10 + pop) overflow -> (ans >= MaxInt32/10)
(ans > MaxInt32/10) -> ans * 10 + pop overflow
(ans == MaxInt32/10) if (pop > 7) then (ans * 10 + pop) overflow
*/
// @lc code=start
func reverse(x int) int {
	ans := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && pop > math.MaxInt32%10) {
			return 0
		}
		if ans < math.MinInt32/10 || (ans == math.MinInt32/10 && pop < math.MinInt32%10) {
			return 0
		}
		ans = ans*10 + pop
	}
	return ans
}

// @lc code=end
