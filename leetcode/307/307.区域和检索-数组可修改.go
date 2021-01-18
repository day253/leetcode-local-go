package leetcode

/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 */

// @lc code=start
// https://leetcode-cn.com/problems/range-sum-query-mutable/solution/gojian-ji-dai-ma-xian-duan-shu-by-niu-meng/
type NumArray struct {
	tree []int
	n    int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := make([]int, 2*n)
	for i, j := n, 0; i < 2*n; i, j = i+1, j+1 {
		tree[i] = nums[j]
	}
	for i := n - 1; i >= 0; i-- {
		tree[i] = tree[2*i] + tree[2*i+1]
	}
	return NumArray{tree, n}
}

func (this *NumArray) Update(index int, val int) {
	index += this.n
	val -= this.tree[index]
	for ; index > 0; index >>= 1 {
		this.tree[index] += val
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	sum := 0
	for left, right = left+this.n, right+this.n; left <= right; {
		if left&1 > 0 {
			sum += this.tree[left]
			left++
		}
		if right&1 <= 0 {
			sum += this.tree[right]
			right--
		}
		left >>= 1
		right >>= 1
	}
	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
// @lc code=end
