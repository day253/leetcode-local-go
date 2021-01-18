package leetcode

import (
	"math"
	"math/rand"
)

/*
 * @lc app=leetcode.cn id=1206 lang=golang
 *
 * [1206] 设计跳表
 */
// https://leetcode-cn.com/problems/design-skiplist/solution/1206-she-ji-tiao-biao-pythonshi-xian-by-tuotuoli/
// @lc code=start
const (
	maxLevel = 16
	maxRand  = 65535.0
)

func randLevel() int {
	return maxLevel - int(math.Log2(1.0+maxRand*rand.Float64()))
}

type skipNode struct {
	value int
	right *skipNode
	down  *skipNode
}

type Skiplist struct {
	head *skipNode
}

func Constructor() Skiplist {
	left := make([]*skipNode, maxLevel)
	right := make([]*skipNode, maxLevel)
	for i := 0; i < maxLevel; i++ {
		left[i] = &skipNode{-1, nil, nil}
		right[i] = &skipNode{200001, nil, nil}
	}
	for i := maxLevel - 2; i >= 0; i-- {
		left[i].right = right[i]
		left[i].down = left[i+1]
		right[i].down = right[i+1]
	}
	left[maxLevel-1].right = right[maxLevel-1]
	return Skiplist{left[0]}
}

func (this *Skiplist) Search(target int) bool {
	node := this.head
	for node != nil {
		if node.right.value > target {
			node = node.down
		} else if node.right.value < target {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

func (this *Skiplist) Add(num int) {
	prev := make([]*skipNode, maxLevel)
	i := 0
	node := this.head
	for node != nil {
		if node.right.value >= num {
			prev[i] = node
			i++
			node = node.down
		} else {
			node = node.right
		}
	}
	n := randLevel()
	arr := make([]*skipNode, n)
	t := &skipNode{-1, nil, nil}
	for i, a := range arr {
		p := prev[maxLevel-n+i]
		a = &skipNode{num, p.right, nil}
		p.right = a
		t.down = a
		t = a
	}
}

func (this *Skiplist) Erase(num int) bool {
	var ans bool
	node := this.head
	for node != nil {
		if node.right.value > num {
			node = node.down
		} else if node.right.value < num {
			node = node.right
		} else {
			ans = true
			node.right = node.right.right
			node = node.down
		}
	}
	return ans
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
// @lc code=end
