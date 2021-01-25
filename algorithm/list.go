package algorithm

// ListNode etc
type ListNode struct {
	Val  int
	Next *ListNode
}

// List2IntArray convert list to []int
func List2IntArray(l *ListNode) []int {
	res := []int{}
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

// IntArray2List convert []int to list
func IntArray2List(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{}
	l := head
	for _, item := range arr {
		l.Next = &ListNode{
			Val: item,
		}
		l = l.Next
	}
	return head.Next
}

// MergeListFunction define MergeListFunction
type MergeListFunction func(*ListNode, *ListNode) *ListNode

// MergeList merge array of lists
// l i arr
// 7   0 1 2 3 4 5 6
// 1 2 |_| |_| |_| |
// 2 4 |___|   |___|
// 4 8 |_______|
// 8   |
func MergeList(listArr []*ListNode, fn MergeListFunction) *ListNode {
	len := len(listArr)

	if len == 0 {
		return nil
	}

	step := 1
	for step <= len {
		for i := 0; i < len; i += step * 2 {
			if i+step < len {
				listArr[i] = fn(listArr[i], listArr[i+step])
			}
		}
		step *= 2
	}

	return listArr[0]
}

// MergeListRecursive return head of two merge list
func MergeListRecursive(l *ListNode, r *ListNode) *ListNode {
	if l == nil {
		return r
	} else if r == nil {
		return l
	} else if l.Val < r.Val {
		l.Next = MergeListRecursive(l.Next, r)
		return l
	} else {
		r.Next = MergeListRecursive(l, r.Next)
		return r
	}
}

// MergeListIterate return head of two merge list
func MergeListIterate(l *ListNode, r *ListNode) *ListNode {
	head := &ListNode{}
	prev := head
	for l != nil && r != nil {
		if l.Val < r.Val {
			prev.Next = l
			l = l.Next
		} else {
			prev.Next = r
			r = r.Next
		}
		prev = prev.Next
	}
	if l != nil {
		prev.Next = l
	} else {
		prev.Next = r
	}
	return head.Next
}
