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
	l := &ListNode{}
	for _, item := range arr {
		l.Next = &ListNode{
			Val: item,
		}
		l = l.Next
	}
	return l.Next
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
	return prev.Next
}
