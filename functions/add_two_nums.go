package functions

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbersV1(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := trans(l1) + trans(l2)
	for l1 = nil; ; {
		l2 = l1
		l1 = &ListNode{Val: int(sum % 10), Next: l2}
		sum = sum / 10
		if sum == 0 {
			break
		}
	}
	// reverse link
	if l1.Next == nil {
		return l1
	}
	if l1.Next.Next == nil {
		l2 = l1.Next
		l1.Next = nil
		l2.Next = l1
		return l2
	}
	var p, l, n *ListNode
	l = l1
	n = l.Next
	for {
		l.Next = p
		p = l
		l = n
		if l == nil {
			return p
		}
		n = n.Next
	}
}

func trans(l1 *ListNode) int64 {
	var sum, n int64
	for n = 1; ; {
		if l1 == nil {
			return sum
		}
		sum += int64(l1.Val) * n
		l1 = l1.Next
		n = n * 10
	}
}

// 暴力法（竟然最快？）
func AddTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	var n0, n1 int
	l0 := &ListNode{}
	resp := l0
	for {
		if l1 != nil && l2 != nil {
			n0 = (l1.Val + l2.Val + n1) % 10
			n1 = (l1.Val + l2.Val + n1) / 10
			l0.Val = n0
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 == nil && l2 != nil {
			n0 = (l2.Val + n1) % 10
			n1 = (l2.Val + n1) / 10
			l0.Val = n0
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			n0 = (l1.Val + n1) % 10
			n1 = (l1.Val + n1) / 10
			l0.Val = n0
			l1 = l1.Next
		}
		if l1 == nil && l2 == nil {
			if n1 != 0 {
				l0.Next = &ListNode{}
				l0.Next.Val = n1
			}
			break
		}
		l0.Next = &ListNode{}
		l0 = l0.Next
	}
	return resp
}
